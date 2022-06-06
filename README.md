# Meditate CLI

Meditate CLI is a command-line tool, created to help you with your meditation/prayer time.

The application allows you to maintain a list of words/concepts that you may want to meditate/pray/reflect on. You can then get it to choose an item from your list for you. The choice is random (unless the list is very small, in which case it is looped through in the order that the items were added). However, the randomness is weighted so that the items that have been chosen recently are less likely to be chosen.

## Installing

If you want to use this application, you can copy the *meditate* executable to somewhere on your machine, and then add that location to your system's PATH variable.

If there is no executable for your OS in this repo, then you can build one using the *makefile*.

## Building the executable

If you have a Go compiler and MAKE installed on your system, then you can build an executable using MAKE. See the *makefile* for more details.

## Commands

Below is the list of commands you can run on meditate:

- ```meditate ``` (With no other arguments) - Chooses and displays an item from your list
- ```meditate add <string>``` - Add a string to your list
- ```meditate remove <string>``` - Remove a string from your list
- ```meditate list``` - Lists the whole list, in alphabetical order
- ```meditate search <string>``` - Search the list for any item containing the given string

[My Twitter: @mattdarbs](http://twitter.com/mattdarbs)  
[My Portfolio](http://md-developer.uk)
