package ip

import (
	"archive/tar"
	"compress/gzip"
	"crypto/md5" //nolint
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

func MoveFile(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		err := inputFile.Close()
		if err != nil {
			return fmt.Errorf("couldn't close input file: %s", err)
		}
		return fmt.Errorf("couldn't open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return err
	}
	err = inputFile.Close()
	if err != nil {
		return err
	}
	if err != nil {
		return fmt.Errorf("writing to output file failed: %s", err)
	}
	// The copy was successful, so now delete the original file
	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("failed removing original file: %s", err)
	}
	return nil
}

// ExtractTarGz extracts a gzipped stream to dest
func ExtractTarGz(r io.Reader, dest string) error {
	zr, err := gzip.NewReader(r)
	if err != nil {
		return fmt.Errorf("stream requires gzip-compressed body: %v", err)
	}

	tr := tar.NewReader(zr)

	for {
		f, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("tar error: %v", err)
		}

		switch f.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(dest+f.Name, 0750); err != nil {
				return fmt.Errorf("extractTarGz: Mkdir() failed: %v", err)
			}
		case tar.TypeReg:
			outFile, err := os.Create(dest + f.Name)
			if err != nil {
				return fmt.Errorf("extractTarGz: Create() failed: %v", err)
			}
			// For our purposes, we don't expect any files larger than 100MiB
			limited := &io.LimitedReader{R: tr, N: 100 << 20}
			if _, err := io.Copy(outFile, limited); err != nil {
				return fmt.Errorf("extractTarGz: Copy() failed: %v", err)
			}
			if err := outFile.Close(); err != nil {
				return err
			}
		default:
			return fmt.Errorf(
				"extractTarGz: %s has uknown type: %v",
				f.Name,
				f.Typeflag)
		}
	}

	return nil
}
func md5Hash(file string) ([]byte, error) {
	filePath := filepath.Clean(file)

	// We know exactly where this file and path is
	// #nosec G304
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	// #nosec G401
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return nil, err
	}

	return h.Sum(nil), f.Close()
}

func VerifyMD5HashFromFile(file, md5sumFile string) error {
	actual, err := md5Hash(file)
	if err != nil {
		return err
	}

	cleanMD5SumFile := filepath.Clean(md5sumFile)

	// We know exactly where this file and path is
	// #nosec G304
	expected, err := ioutil.ReadFile(cleanMD5SumFile)
	if err != nil {
		return err
	}

	if fmt.Sprintf("%x", actual) != fmt.Sprintf("%s", expected) {
		return errors.New("checksum error")
	}

	return nil
}

// FindFile returns a path to a file matching regex under root
// Returns
//
//	string: Full path
//	string: File name
//	error : Error
func FindFile(root, r string) (string, string, error) {
	regex, err := regexp.Compile(r)
	if err != nil {
		return "", "", err
	}

	var foundPath string
	var foundName string

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if regex.MatchString(info.Name()) {
			foundPath = path
			foundName = info.Name()
		}
		return nil
	})

	if err != nil {
		return "", "", err
	}
	return foundPath, foundName, nil
}
