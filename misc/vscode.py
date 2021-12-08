import os
import shutil
from . import topology


class VSCodeConfig:

    def __init__(self, top, dry_run=False):
        self.top = top
        self.dry_run = dry_run

    
    def save(self, repo):

        r = self.top.get(repo)
        scripts = os.path.join(self.top.scripts(), "vscode-configs", repo)

        to_save = os.path.join(r, ".vscode")
        if not os.path.exists(to_save):
            print("No config to save")
            return

        if not self.dry_run:
            shutil.copytree(to_save, scripts)
        print(f"Saved {to_save} in {scripts}")

    def restore(self, repo):
        in_repo = os.path.join(self.top.get(repo), ".vscode")
        to_restore = os.path.join(self.top.scripts(), "vscode-configs", repo)

        if not os.path.exists(to_restore):
            print("No config to save")
            return

        if not self.dry_run:
            shutil.copytree(to_restore, in_repo)
        print(f"Saved {to_restore} in {in_repo}")
