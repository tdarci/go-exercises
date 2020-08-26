# Git FTW

What is git?

Why should I use git?

How do I use git?

## What Is Git?

You have a lot of files in your codebase.

    Git is a thing

        that allows you

           to keep snapshots of the complete state of these files.

Meaning, that you can...

- ...make changes to a bunch of files and keep track of all these changes as one thing (sort of a "snapshot" of the state of all your files).
- ...make changes to a bunch of files at the same time somebody else is also making changes.
- ...work on a feature and keep it contained inside one git thing while working on another feature inside another git thing.

Simply put, git is a version control system. An elegant, powerful, surprisingly lightweight, distributed version control system.

### Ah, Is It the Same as "GitHub"?

No.

Git is an executable program that you run on your machine. Github is a service that is built _on top of git_.

There are two big benefits of github:

- Github gives you a place in the ☁️ to store a copy of your git repository. (And therefore a place for several people to easily share their code.)
- Github has a helpful system for reviewing code before it becomes part of the main area of the code (a.k.a. "pull requests").

Because of these benefits, __I recommend using gitHub as part of your git solution__.

## Why Should I Use Git (and GitHub)?

If you are working alone on a project, the benefits of git are not as significant as if you are working with other people, where using git is a no-brainer.

Even when working alone, the ability to work on a feature as a contained unit and then decide when to add it to your codebase... and even to be able to roll back the whole feature if you later find it is problematic... is very helpful. And, of course, having a place other than your laptop where your code is stored is essential. (And the free price for small projects makes the choice much easier.)

In addition, the world uses git and gitHub. Being well-versed in git will help you get up to speed on a development team. This includes the world of open-source.

If you are working with others, the ability to have a centralized repository of your code for all to work on is a big benefit, and the code review process of gitHub is a de-facto industry standard of some parts of the software world.

What about Mercurial? Or Subversion, CSV, Perforce, Team Foundation Server, Darcs, or SourceSafe? Using _any_ version control system will introduce you to the benefits of working this way. Git's ease of branching, coupled with its wide use in software development, make it the choice for so many, including my workplaces for the past several jobs. That said, it is not without its headaches and learning curve. And I would not be the person I am today had I not struggled with Visual SourceSafe 6.0.

## How Do I Use Git?

There is a lot you can do with git. What follows is a path through the commands you will need to get a good working process. Also note that gitHub aims to make some of git easier via UI buttons, but my tendency is to stay with the command line for most things.

### Get Set Up

1. Install git.
2. Create a gitHub account (and setup ssh access to it).
3. As a helper to see what's going on (but NOT to run git commands), you may want to install one of these: SourceTree from Atlassian (Mac, Windows), GitKraken (Linux). [See here for more](https://git-scm.com/downloads/guis).

### Create a Repository

This is something you likely do just once in a while, but it is an important first step in getting going with a set of code.

A "repository" (a.k.a. "repo") is a container for all of the code on a single project in git. There are philosophies about how best to define the scope of a git repo (from small projects to monorepos). But for the purposes of what we're up to here, I recommend creating one repository for all of your work on Go exercises, a place for all of the code you're making when learning go. It can contain a directory for each exercise as a way to break it up.

To create a repository:

1. Create the directory.
2. Initialize it with git.
3. Change the name of the default branch.
4. Add it to gitHub. Make sure the default branch is set.
5. Make your first commit, and see that it's all working.

#### Create a Repository: Create the Directory

#### Create a Repository: Initialize it with Git

#### Create a Repository: Change the Name of the Default Branch

#### Create a Repository: Add it to GitHub

#### Create a Repository: Make Your First Commit

### Work on a Feature

This is the process:

1. Create a branch.
2. Work... (do this until the code is ready for review)
    1. Write code.
    2. Make a commit.
    3. Push commit to gitHub.
3. Submit a pull request for your branch.
4. Request reviews on your pull request.
5. Adjust code based on feedback (see step 2)
6. Merge your pull request into the default branch (and delete the remote and local branches)

## Bonus: How Does Git Work?

Once you've used git for a little while, read [Git from the Bottom Up](https://jwiegley.github.io/git-from-the-bottom-up/). It's illuminating.