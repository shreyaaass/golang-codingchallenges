# Problem 1

## File Information CLI

This Go program is a command-line tool for obtaining various statistics about a given file. It provides options to count the number of bytes, lines, words, and characters in a file.

### Command-line Flags

- `-c`: Enables counting the number of bytes in a file.
- `-l`: Enables counting the number of lines in a file.
- `-w`: Enables counting the number of words in a file.
- `-m`: Enables counting the number of characters in a file.
- `-f [file_path]`: Specifies the path to the file. If not provided, it defaults to "test.txt".

### Main Function

The `main` function begins by defining flags for each statistical option and the file path. It then parses the command-line flags provided by the user.

### Reading File Data

The program reads the contents of the file specified by the user using `os.ReadFile(*file_path)`. If an error occurs during file reading, it prints an error message and exits.

### Flag Handling

- If no flags are provided (`flag.NFlag()==0`), the program sets all statistical options to true by default.
- Based on the enabled flags, the program determines whether to count bytes, lines, words, or characters in the file.
- If the `-c` flag is enabled, the program counts the number of bytes in the file.
- If the `-m` flag is enabled, the program counts the number of characters using `utf8.RuneCountInString`.
- If the `-l` flag is enabled, the program counts the number of lines using `strings.Count`.
- If the `-w` flag is enabled, the program counts the number of words using `strings.Count`.

### Output

The program prints the counted statistics for the file in the following format:
'''[Number of bytes] [Number of lines] [Number of words] [File path]'''