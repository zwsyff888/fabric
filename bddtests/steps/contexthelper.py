
import uuid
import os
import shutil
from slugify import slugify

class ContextHelper:

    @classmethod
    def GetHelper(cls, context):
        if not "contextHelper" in context:
            context.contextHelper = ContextHelper(context)
        return context.contextHelper

    def __init__(self, context):
        self.context = context
        self.guuid = str(uuid.uuid1()).replace('-','')

    def getBootrapHelper(self, chainId):
        import bootstrap_util
        return bootstrap_util.BootstrapHelper(chainId=chainId)

    def getGuuid(self):
        return self.guuid

    def getTmpPath(self):
        pathToReturn = "tmp"
        if not os.path.isdir(pathToReturn):
            os.makedirs(pathToReturn)
        return pathToReturn

    def getCachePath(self):
        pathToReturn = os.path.join(self.getTmpPath(), "cache")
        if not os.path.isdir(pathToReturn):
            os.makedirs(pathToReturn)
        return pathToReturn


    def getTmpProjectPath(self):
        pathToReturn = os.path.join(self.getTmpPath(), self.guuid)
        if not os.path.isdir(pathToReturn):
            os.makedirs(pathToReturn)
        return pathToReturn

    def getTmpPathForName(self, name, extension=None, copyFromCache=False):
        'Returns the tmp path for a file, and a flag indicating if the file exists. Will also check in the cache and copy to tmp if copyFromCache==True'
        slugifiedName = ".".join([slugify(name), extension]) if extension else slugify(name)
        tmpPath = os.path.join(self.getTmpProjectPath(), slugifiedName)
        fileExists = False
        if os.path.isfile(tmpPath):
            # file already exists in tmp path, return path and exists flag
            fileExists = True
        elif copyFromCache:
            # See if the file exists in cache, and copy over to project folder.
            cacheFilePath = os.path.join(self.getCachePath(), slugifiedName)
            if os.path.isfile(cacheFilePath):
                shutil.copy(cacheFilePath, tmpPath)
                fileExists = True
        return (tmpPath, fileExists)

    def copyToCache(self, name):
        srcPath, fileExists = self.getTmpPathForName(name, copyFromCache=False)
        assert fileExists, "Tried to copy source file to cache, but file not found for: {0}".format(srcPath)
        # Now copy to the cache if it does not already exist
        cacheFilePath = os.path.join(self.getCachePath(), slugify(name))
        if not os.path.isfile(cacheFilePath):
            shutil.copy(srcPath, cacheFilePath)


    def isConfigEnabled(self, configName):
        return self.context.config.userdata.get(configName, "false") == "true"

    def before_scenario(self, scenario):
        print("before_scenario: {0}".format(self))

    def after_scenario(self, scenario):
        print("after_scenario: {0}".format(self))


    def before_step(self, step):
        print("before_step: {0}".format(self))
        print("")

    def registerComposition(self, composition):
        return composition

