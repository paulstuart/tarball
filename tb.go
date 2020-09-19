// Package tarball creates and reads from tar files
package tarball

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Create creates a tar globbing src
func Create(w io.Writer, srcs ...string) error {
	tw := tar.NewWriter(w)
	for _, src := range srcs {
		files, err := filepath.Glob(src)
		if err != nil {
			return err
		}
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				return err
			}
			stat, err := f.Stat()
			if err != nil {
				return err
			}
			hdr := &tar.Header{
				Name: file,
				Mode: int64(stat.Mode()),
				Size: stat.Size(),
			}
			if err := tw.WriteHeader(hdr); err != nil {
				return err
			}
			//if _, err := tw.Write([]byte(file.Body)); err != nil {
			if _, err = io.Copy(tw, f); err != nil {
				return err
			}
		}
	}
	return tw.Close()
}

// ReadAll dumps the output of the files
// Primarily for testing
func ReadAll(r io.Reader, w io.Writer) error {
	// Open and iterate through the files in the archive.
	tr := tar.NewReader(r)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "Contents of %s:\n", hdr.Name)
		if _, err := io.Copy(w, tr); err != nil {
			return err
		}
		fmt.Fprintln(w)
	}
	return nil
}

// ReadFile extracts the file from the tar file
func ReadFile(r io.Reader, w io.Writer, filename string) error {
	// Open and iterate through the files in the archive.
	tr := tar.NewReader(r)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			return err
		}
		if filename != hdr.Name {
			continue
		}
		_, err = io.Copy(w, tr)
		return err
	}
	return fmt.Errorf("file %q not found", filename)
}

// FileList returns the files in the archive
func FileList(r io.Reader) ([]string, error) {
	var list []string
	tr := tar.NewReader(r)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			return nil, err
		}
		list = append(list, hdr.Name)
	}
	return list, nil
}
