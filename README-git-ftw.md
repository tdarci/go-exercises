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

Git is an executable program that you run on your machine. Github is a service
that is built _on top of git_.

There are two big benefits of github:

- Github gives you a place in the ☁️ to store a copy of your git repository. (And therefore a place for several people to easily share their code.)
- Github has a helpful system for reviewing code before it becomes part of the main area of the code (a.k.a. "pull requests").

Because of these benefits, __I recommend using gitHub as part of your git solution__.

## Why Should I Use Git (and GitHub)?

If you are working alone on a project, the benefits of git are not as
significant as if you are working with other people, where using git is a
no-brainer.

Even when working alone, the ability to work on a feature as a contained unit
and then decide when to add it to your codebase... and even to be able to roll
back the whole feature if you later find it is problematic... is very helpful.
And, of course, having a place other than your laptop where your code is
stored is essential. (And the free price for small projects makes the choice
much easier.)

In addition, the world uses git and gitHub. Being well-versed in git will help
you get up to speed on a development team. This includes the world of
open-source.

If you are working with others, the ability to have a centralized repository
of your code for all to work on is a big benefit, and the code review process
of gitHub is a de-facto industry standard of some parts of the software world.

What about Mercurial? Or Subversion, CSV, Perforce, Team Foundation Server,
Darcs, or SourceSafe? Using _any_ version control system will introduce you to
the benefits of working this way. Git's ease of branching, coupled with its
wide use in software development, make it the choice for so many, including my
workplaces for the past several jobs. That said, it is not without its
headaches and learning curve. And I would not be the person I am today had I
not struggled with Visual SourceSafe 6.0.

## How Do I Use Git?

There is a lot you can do with git. What follows is a path through the
commands you will need to get a good working process. Also note that gitHub
aims to make some of git easier via UI buttons, but my tendency is to stay
with the command line for most things.

But first, these resources may be helpful (now or to come back to for reference):

- [Git tutorial for beginners](https://www.code-learner.com/git-tutorial-for-beginners/)
- [This visual walk-through can help you see some of the git concepts](https://agripongit.vincenttunru.com/)
- [This visual guide is kind of interesting too](https://marklodato.github.io/visual-git-guide/index-en.html).
- [The docs](https://git-scm.com/docs)
- [This interactive thing is cool](https://ndpsoftware.com/git-cheatsheet.html)

------------------------------------------------------------

### Get Set Up

1. Install git.
2. Create a gitHub account (and setup ssh access to it).
3. As a helper to see what's going on (but NOT to run git commands), you may want to install one of these: SourceTree from Atlassian (Mac, Windows), GitKraken (Linux). [See here for more](https://git-scm.com/downloads/guis).

------------------------------------------------------------

### Create a Repository

This is something you likely do just once in a while, but it is an important
first step in getting going with a set of code.

A "repository" (a.k.a. "repo") is a container for all of the code on a single
project in git. There are philosophies about how best to define the scope of a
git repo (from small projects to monorepos). But for the purposes of what
we're up to here, I recommend creating one repository for all of your work on
Go exercises, a place for all of the code you're making when learning go. It
can contain a directory for each exercise as a way to break it up.

To create a repository:

1. Create the directory.
2. Initialize it with git.
3. Change the name of the default branch.
4. Create your first commit.
5. Add it to gitHub. Make sure the default branch is set.

#### Create a Repository: Create the Directory

Change to the directory where you would like to store this code. Perhaps `cd
~/src`

Create the directory for your exercises. Perhaps `mkdir my-go`

Go there: `cd my-go`

#### Create a Repository: Initialize it with Git

You are in the empty directory you just created. Make it a git directory by
running `git-init`.

#### Create a Repository: Change the Name of the Default Branch

In your newly-initialized git directory, create a branch that will be your
default branch: `git checkout -b main`

#### Create a Repository: Create Your First Commit

```
echo "# my-go" > README.md
git add README.md
git commit -m 'first commit'
```

#### Create a Repository: Add it to GitHub

1. on gitHub...
    1. Login to gitHub.
    2. Press the "+" button and select "New Repository"
    3. Enter the repository name "my-go". You can make it public or private. Do **not** initialize the repository with a README file or anything else.
2. in terminal... (connect your local reposotitory to github)
    1. `git remote add origin git@github.com:YOUR_GITHUB_USERNAME/my-go.git`
    2. `git push origin main`

Now you have a repository you can work in and that you can sync to its remote
counterpart on gitHub.

------------------------------------------------------------

### Work on a Feature

Now that you have a repository, you will be adding code to it over the next
days, weeks and maybe years.

Each time you are adding a new feature, or fixing an existing feature, you
will go through the process described here. So often that it will become
second nature. The basic idea is to create a "branch", work in it until your
code is ready, and then merge the code into `main`.

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

#### Work on a Feature: Create a branch

Let's say we're working on a new image filter...

First, be sure you are on the `main` branch: `git checkout main`.

And be sure you have no local changes. When you type `git status` you should see:

```
On branch main
nothing to commit, working tree clean
```

Create your branch: `git checkout -b cats-filter`

#### Work on a Feature: Work...

Edit and create files...

Type `git status` to see what you have:

```
On branch cats-filter
Untracked files:
  (use "git add <file>..." to include in what will be committed)

    filter-exercise/

nothing added to commit but untracked files present (use "git add" to track)
```

This shows me that I have a new directory in my workspace. Let's add
everything in it to my "index":

`git add .`

And now `git status` shows me:

```
On branch cats-filter
Changes to be committed:
  (use "git reset HEAD <file>..." to unstage)

    new file:   filter-exercise/filters/cats.go
    new file:   filter-exercise/filters/filter.go
    new file:   filter-exercise/main.go
```

These files are ready to be added to a "commit". A commit is a snapshot of all
the files in your repository.

Commit the files like so: `git commit -m 'Create cats filter'`

We now see this for `git status`:

```
On branch cats-filter
nothing to commit, working tree clean
```

We can push these changes up to gitHub like so (into a new, remote copy of the
`cats-filter` branch): `git push --set-upstream origin cats-filter`

Note that after this first push for the branch, subsequent commits may be
pushed to the branch with just: `git push`.

#### Work on a Feature: Submit a PR

When your code is ready to have somebody look at it, create a "pull request"
(also known as a "PR") in gitHub.

You can do this by going to the main page for your repo, where you should see
a button for creating a pull request for your branch (called "Compare & pull
request"), or you can just navigate to
`https://github.com/YOUR_GITHUB_USERNAME/REPO_NAME/pull/new/BRANCH_NAME`.

#### Work on a Feature: Request Reviews

On the page for your PR you can enter a title and description, and you can
select reviewers. Comments your reviewers enter will appear here for you to
see and respond to.

#### Work on a Feature: Adjust Code

Make adjustments to your code based on what you learn in the code review. Use
the same process for committing code as described in the "Work..." section.

#### Work on a Feature: Merge Branch into Default Branch

When you are happy with the state of your new code, press the `Merge` button
on gitHub to merge it into the `main` branch.

------------------------------------------------------------

### Git Commands to Know

// TODO: Add diagram of... local... remote ("origin")... github... git runs
// where... local commands vs. commands with remotes... peer to peer...

// TODO: Add starter set of git terms. repository. commit. workspace. index.
// branch. tag. remote. origin. .git directory. 

#### clone: Get a copy of a repository from any git server onto your local machine.

example: `git clone git@github.com:tdarci/mygo.git`

#### fetch: Get the latest changes that exist on your main remote git server(a.k.a. "origin"), without changing any of your local code.

`git fetch` grabs all the latest changes from the "origin" remote and brings them to your machine.

It **does not** overwrite anything you are working on. Instead, it creates
"shadow branches" of your existing branches that start with
`/remotes/origin/`. Here is an example `git branch -a`:

```
* cats-filter
  main
  remotes/origin/cats-filter
  remotes/origin/main
```

#### checkout

Checkout does a lot. Too much maybe.

Create a new branch like so, based on the current commit you are on: 
`git checkout -b my-new-branch-name`

Checkout also lets you grab versions of individual files or kind of destroy
your current workspace.

#### branch

Run `git branch` or `git branch -a` to see the branches on your machine for
the current repo.

#### add

Use `git add` to add files to the "index" (where they are now ready to be
committed). Some examples:

- `git add .`: Add all the changed and new files that are here to the index.
- `git add -u`: Add all the changed files to the index.

#### commit

Run `git commit -m 'descritiption of commit'` to create a git commit. The
commit is the basic unit of work in the world of git.

#### push

Push the current branch to the remote origin.

For a branch that has never been sent to origin: `git push --set-upstream origin my-branch-name`.

For a branch that has already been sent: `git push`

#### merge

Use `git merge` to bring code together. Here is an example...

```
# =========== Start by getting in-sync with main:
git fetch

git checkout main

git status # to check that there are no local changes and we are up to date

# if we are behind:
git rebase

# =========== Then, do some work...
git checkout -b new-branch

...make some changes... 

git add .
git commit -m 'cool things'

# =========== OK, let's bring these changes back into main.
git checkout main

git merge new-branch

git push

```

#### rebase

Rebase is a weird and important command.

Each commit in git knows its parent.

Rebase lets you take an existing commit and create a new commit that looks
just like it but has a _different parent_. This can be very helpful when
working on a branch that is open for some time, and while you're working,
`main` is also changing. With `rebase` you can make your branch code be a
direct descendant of the current `master`, instead of getting behind.

## Bonus: How Does Git Work?

__Mind-Bender (a must-read):__ Once you've used git for a little while, read [Git from the Bottom Up](https://jwiegley.github.io/git-from-the-bottom-up/). __It's illuminating.__
