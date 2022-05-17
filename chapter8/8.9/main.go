package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

type sizeByRoot struct {
	root string
	size int64
}

func main() {
	flag.Parse()

	// Determine the initial directories.
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse each root of the file tree in parallel.
	fileSizes := make(chan sizeByRoot)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// Print the results periodically.
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}

	nfiles := make(map[string]int64)
	nbytes := make(map[string]int64)

loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles[size.root]++
			nbytes[size.root] += size.size
		case <-tick:
			printDisUsageByRoots(nfiles, nbytes)
		}
	}

	printDiskUsageTotal(nfiles, nbytes) // final totals
}

func printDiskUsageTotal(nfiles, nbytes map[string]int64) {
	var totalFiles, totalSize int64
	for root := range nfiles {
		totalFiles += nfiles[root]
		totalSize += nbytes[root]
	}

	fmt.Printf("Total %d files  %.1f GB\n", totalFiles, float64(totalSize)/1e9)
}

func printDisUsageByRoots(nfiles, nbytes map[string]int64) {
	for root := range nfiles {
		fmt.Printf("root %s: %d files  %.1f GB\n", root, nfiles[root], float64(nbytes[root])/1e9)
	}
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(originalDir string, dir string, n *sync.WaitGroup, fileSizes chan<- sizeByRoot) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(originalDir, subdir, n, fileSizes)
		} else {
			fileSizes <- sizeByRoot{root: originalDir, size: entry.Size()}
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
