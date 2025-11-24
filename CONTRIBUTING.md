## A brief guide on contributing to the repository

### 1. Clone the Repository

If you haven't already, clone the repository to your local machine.

```Bash
git clone git@github.com:WGU-Independent-Student-Group/Webapp.git
cd Webapp
```

### 2. Update Main

Before starting new work, ensure your local main is up to date.

```Bash
git checkout main
git pull origin main
```

### 3. Create a New Branch

Never work directly on main. Create a descriptive branch for your changes. A common convention is: type/description (e.g., feature/login-page or fix/header-bug)

```Bash
git checkout -b feature/my-new-feature
```

### 4. Make Changes and Commit

Do all of your work in that branch. Once satisfied, stage and commit your changes.

```Bash
# Staging
git add .
```

Commit with a clear message

```Bash
# Commiting
git commit -m "Add brief description of what changed"
```

### 5. Push Your Branch

Push your branch to the remote repository (origin).

```Bash
git push -u origin feature/my-new-feature
```

### 6. Submit a Pull Request (PR)

Go to the repository on GitHub.

Click the Pull requests tab.

Click New pull request.

Select your branch in the "compare" dropdown.

Fill out the title and description.

Click Create pull request.

What Happens Next?

Team members will review your code and may request changes.

Once approved, a maintainer will merge your branch into main.

After the merge, you can delete your feature branch.
