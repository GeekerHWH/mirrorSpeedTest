# A project for Testing debian mirrors
This tool is mainly to help Debian users to filter the best apt Repo for themselves,
especially for people who live in China with a special Network Environment.

**We geekers always want to be the fastest. right?**

If you have any suggestion or bug problem feel free to address an issue:)

# What's next
- [x] solved contentLength = -1 problem so that the speed won't be negative
- [x] fixed a bug that cause tool crash
- [x] input multiple URLs to test speeds
- [x] added changeMirror.sh to change the mirror supported by default
- [x] modulize speed test code
- [x] multi choose mirrors to test
- [ ] write a chinese README.md
- [ ] support English Comments
- [ ] support English version
- [ ] check whether the host OS is Debian
- [ ] support more mirrors
- [ ] support multi-threads testing in parallel
- [ ] support geo-based mirrors testing to present perfect suggestion
- [ ] support delay testing
- [ ] support integrated apt sourcelist file editting(default choose the fastest)
- [ ] beyond my imagination...

# How to use it?
1. make sure your computer has the Go runtime installed
```bash
go env
```
2. enter the directory of this project, then open the terminal to run:
```bash
go run main/main.go
```
3. follow the instruction of the app to get your best mirror(English version 
coming soon)