# Hello World (Go)

## Purpose
The purpose of this project was to learn to hook my Atom Editor to both Go and Github. The definition of done is as followed:
- [x] Connect git to Github
- [x] Run hello world go program  

## How to connect to git
To connect your Atom browser to get, first watch some youtube videos on how to set your Atom IDE up to theoretically connect to git. Then (honestly this was a moth or two ago so I really don't remember that well how to get this done... sorry friends.) once that is setup, you can edit your .git/congif file to include the following lines

```
[remote "origin"]
			url = "https://github.com/mjschmidt/hello.git"

[branch "master"]
	remote = origin
	merge = refs/heads/master
```

After that you will need to create the git project on github so that your project has something to connect up to. But once that is done, you should be able to do git things with your git bash from online connected up to your github account. Again the bottom left has the ability to fetch, pull and push. Left click or setup your hotkeys to be able to do this automagically.

## Running go programs,
So you will want to setup your Atom IDE to be able to run your computer's command line stuff. To do this install a terminal package through the Atom, download packages area. Right now to bring that up you can do a ctr + sift + p and type install packages.


Once you have your terminal installed you can use the hotkey of ctr +  ` to open your shell session. and from there its is a simple

```
go run hello.go
```

```
PS C:\Users\Mike\Documents\GitHub\Atom\src\hello> go run hello.go
hello world
PS C:\Users\Mike\Documents\GitHub\Atom\src\hello>
```
