package app

func (fl *Service) Start(module, basePath string) {
	fl.CreateAllFolders()
	fl.CreateAllFiles(module, basePath)
}
