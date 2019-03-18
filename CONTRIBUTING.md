# Contributing

<!-- SPDX-License-Identifier: (MIT OR CC-BY-3.0+) -->

Feedback and contributions are very welcome!

Here's help on how to make contributions, divided into the following sections:

* general information,
* vulnerability reporting,
* documentation changes,
* code changes,
* how to check proposed changes before submitting them,
* reuse (supply chain for third-party components, including updating them), and
* keeping up with external changes.

## General information

For specific proposals, please provide them as
[pull requests](https://github.com/coreinfrastructure/best-practices-badge/pulls)
or
[issues](https://github.com/coreinfrastructure/best-practices-badge/issues)
via our
[GitHub site](https://github.com/coreinfrastructure/best-practices-badge).

The "doc/" directory has information you may find helpful, for example:

The [INSTALL.md](doc/INSTALL.md) file explains how to install the program
locally (highly recommended if you're going to make code changes).
It also provides a quick start guide.

### Pull requests and different branches recommended

Pull requests are preferred, since they are specific.
For more about how to create a pull request, see
<https://help.github.com/articles/using-pull-requests/>.

We recommend creating different branches for different (logical)
changes, and creating a pull request when you're done into the master branch.
See the GitHub documentation on
[creating branches](https://help.github.com/articles/creating-and-deleting-branches-within-your-repository/)
and
[using pull requests](https://help.github.com/articles/using-pull-requests/).

### How we handle proposals

We use GitHub to track proposed changes via its
[issue tracker](https://github.com/Labbs/gitlab-ci-pipelines-exporter/issues) and
[pull requests](https://github.com/Labbs/gitlab-ci-pipelines-exporter/pulls).
Specific changes are proposed using those mechanisms.
Issues are assigned to an individual, who works it and then marks it complete.
If there are questions or objections, the conversation area of that
issue or pull request is used to resolve it.

### Two-person review

Our policy is that at least 50% of all proposed modifications will be reviewed
before release by a person other than the author,
to determine if it is a worthwhile modification and free of known issues
which would argue against its inclusion
(per the Gold requirement two_person_review).

We achieve this by splitting proposals into two kinds:

1. Low-risk modifications.  These modifications are being proposed by
   people authorized to commit directly, pass all tests, and are unlikely
   to have problems.  These include documentation/text updates
   (other than changes to the criteria) and/or updates to existing gems
   (especially minor updates) where no risk (such as a security risk)
   have been identified.  The project lead can decide that any particular
   modification is low-risk.
2. Other modifications.  These other modifications need to be
   reviewed by someone else or the project lead can decide to accept
   the modification.  Typically this is done by creating a branch and a
   pull request so that it can be reviewed before accepting it.

### Developer Certificate of Origin (DCO)

All contributions (including pull requests) must agree to
the [Developer Certificate of Origin (DCO) version 1.1](doc/dco.txt).
This is exactly the same one created and used by the Linux kernel developers
and posted on <http://developercertificate.org/>.
This is a developer's certification that he or she has the right to
submit the patch for inclusion into the project.

Simply submitting a contribution implies this agreement, however,
please include a "Signed-off-by" tag in every patch
(this tag is a conventional way to confirm that you agree to the DCO).
You can do this with <tt>git commit --signoff</tt> (the <tt>-s</tt> flag
is a synonym for <tt>--signoff</tt>).

Another way to do this is to write the following at the end of the commit
message, on a line by itself separated by a blank line from the body of
the commit:

````
Signed-off-by: YOUR NAME <YOUR.EMAIL@EXAMPLE.COM>
````

You can signoff by default in this project by creating a file
(say "git-template") that contains
some blank lines and the signed-off-by text above;
then configure git to use that as a commit template.  For example:

````sh
git config commit.template ~/cii-best-practices-badge/git-template
````

It's not practical to fix old contributions in git, so if one is forgotten,
do not try to fix them.  We presume that if someone sometimes used a DCO,
a commit without a DCO is an accident and the DCO still applies.

### License (MIT)

All (new) contributed material must be released
under the [MIT license](./LICENSE).
All new contributed material
that is not executable, including all text when not executed,
is also released under the
[Creative Commons Attribution 3.0 International (CC BY 3.0) license](https://creativecommons.org/licenses/by/3.0/) or later.

See the section on reuse for their license requirements
(they don't need to be MIT, but all required components must be
open source software).

### No trailing whitespace

Please do not use or include trailing whitespace
(spaces or tabs at the end of a line).
Since they are often not visible, they can cause silent problems
and misleading unexpected changes.
For example, some editors (e.g., Atom) quietly delete them by default.

## Vulnerability reporting (security issues)

If you find a significant vulnerability, or evidence of one,
please send an email to the security contacts that you have such
information, and we'll tell you the next steps.
For now, the security contacts are:
Jeremy Mouton <moutonjeremy@labbs.fr>.

Please use an email system (like Gmail) that supports
hop-to-hop encryption using STARTTLS when reporting vulnerabilities.
Examples of such systems include Gmail, Outlook.com, and runbox.com.
See [STARTTLS Everywhere](https://starttls-everywhere.org/)
if you wish to learn more about efforts to encourage the use of STARTTLS.
Your email client should use encryption to communicate with
your email system (i.e., if you use a web-based email client then use HTTPS,
and if you use email client software then configure it to use encryption).
Hop-to-hop encryption isn't as strong as end-to-end encryption,
but we've decided that it's strong enough for this purpose
and it's much easier to get everyone to use it.

We will gladly give credit to anyone who reports a vulnerability
so that we can fix it.
If you want to remain anonymous or pseudonymous instead,
please let us know that; we will gladly respect your wishes.

## Documentation changes

Most of the documentation is in "markdown" format.
All markdown files use the .md filename extension.

Where reasonable, limit yourself to Markdown
that will be accepted by different markdown processors
(e.g., what is specified by CommonMark or the original Markdown)
In practice we use
the version of Markdown implemented by GitHub when it renders .md files,
and you can use its extensions
(in particular, mark code snippets with the programming language used).
This version of markdown is sometimes called
[GitHub-flavored markdown](https://help.github.com/articles/github-flavored-markdown/).
In particular, blank lines separate paragraphs; newlines inside a paragraph
do *not* force a line break.
Beware - this is *not*
the same markdown algorithm used by GitHub when it renders
issue or pull comments; in those cases
[newlines in paragraph-like content are considered as real line breaks](https://help.github.com/articles/writing-on-github/);
unfortunately this other algorithm is *also* called
GitHub rendered markdown.
(Yes, it'd be better if there were standard different names
for different things.)

The style is basically that enforced by the "markdownlint" tool.
Don't use tab characters, avoid "bare" URLs (in a hypertext link, the
link text and URL should be on the same line), and try to limit
lines to 80 characters (but ignore the 80-character limit if that would
create bare URLs).
Using the "rake markdownlint" or "rake" command
(described below) implemented in the development
environment can detect some problems in the markdown.
That said, if you don't know how to install the development environment,
don't worry - we'd rather have your proposals, even if you don't know how to
check them that way.

Do not use trailing two spaces for line breaks, since these cannot be
seen and may be silently removed by some tools.
Instead, use <tt>&lt;br&nbsp;/&gt;</tt> (an HTML break).

## Code changes

To make changes to the "BadgeApp" web application that implements the criteria,
you may find the following helpful; [INSTALL.md](doc/INSTALL.md)
(installation information) and [implementation.md](doc/implementation.md)
(implementation information).

The code should strive to be DRY (don't repeat yourself),
clear, and obviously correct.
Some technical debt is inevitable, just don't bankrupt us with it.
Improved refactorizations are welcome.

### Testing during continuous integration

Note that we also use
[CircleCI](https://circleci.com/gh/coreinfrastructure/best-practices-badge)
for continuous integration tools to check changes
after they are checked into GitHub; if they find problems, please fix them.
These run essentially the same set of checks as the default rake task.

## Git commit messages

When writing git commit messages, try to follow the guidelines in
[How to Write a Git Commit Message](https://chris.beams.io/posts/git-commit/):

1.  Separate subject from body with a blank line
2.  Limit the subject line to 50 characters.
    (We're flexible on this, but *do* limit it to 72 characters or less.)
3.  Capitalize the subject line
4.  Do not end the subject line with a period
5.  Use the imperative mood in the subject line (*command* form)
6.  Wrap the body at 72 characters ("<tt>fmt -w 72</tt>")
7.  Use the body to explain what and why vs. how
    (git tracks how it was changed in detail, don't repeat that)

## Reuse (supply chain)

### Requirements for reused components

We prefer reusing components instead of writing lots of code,
but please evaluate all new components before adding them
(including whether or not you need them).
We want to reduce our risks of depending on software that is poorly
maintained or has vulnerabilities (intentional or unintentional).

Mike Perham's [Kill Your Dependencies](http://www.mikeperham.com/2016/02/09/kill-your-dependencies/)
notes that, "every dependency in your application has the potential to
bloat your app, to destabilize your app, to inject odd behavior...
When you are considering adding a dependency to your Rails app,
it's a good idea to do a quick sanity check...".
So don't bring in gems you don't need
(if it's trivial to re-implement the required function, consider doing it).
Also, if the gem transitively depends on in many other gems,
especially if they are new additions, look for simpler alternatives
or help the upstream library remove the unnecessary dependencies.

## Keeping up with external changes

The installer adds a git remote named 'upstream'.
Running 'git pull upstream master' will pull the current version from
upstream, enabling you to sync with upstream.

You can reset this, if something has happened to it, using:

~~~~sh
git remote add upstream \
    https://github.com/Labbs/gitlab-ci-pipelines-exporter.git
~~~~

If the version of Ruby has changed (in the Gemfile),
use the 'Ruby itself can be updated' instructions.
If gems have been added or their versions changed,
run "bundle install" to install the new ones.
