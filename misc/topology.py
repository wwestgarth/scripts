import os 


class Repos:

    def __init__(self):
        pass

    def get(self, repo):
        if repo in self.repos():
            return os.path.join(self.work(), repo)
        return None

    @staticmethod
    def default_vegahome():
        return "/Users/wwestgarth/wort/vegahome"

    @staticmethod
    def scripts():
        return os.path.join(Repos.work(), "scripts")

    @staticmethod
    def work():
        return "/Users/wwestgarth/work"

    @staticmethod
    def repos():
        return {"vega", "data-node", "protos", "vegawallet"}

    @staticmethod
    def current():
        cwd = os.getcwd()

        for repo in Repos.repos():
            if cwd == os.path.join(Repos.work(), repo):
                return repo
        return None
