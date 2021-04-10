package build

import (
	"github.com/docker-slim/docker-slim/pkg/app/master/commands"

	"github.com/c-bata/go-prompt"
)

var CommandSuggestion = prompt.Suggest{
	Text:        Name,
	Description: Usage,
}

var CommandFlagSuggestions = &commands.FlagSuggestions{
	Names: []prompt.Suggest{
		{Text: commands.FullFlagName(commands.FlagTarget), Description: commands.FlagTargetUsage},
		{Text: commands.FullFlagName(commands.FlagPull), Description: commands.FlagPullUsage},
		{Text: commands.FullFlagName(commands.FlagShowPullLogs), Description: commands.FlagShowPullLogsUsage},
		{Text: commands.FullFlagName(FlagShowBuildLogs), Description: FlagShowBuildLogsUsage},
		{Text: commands.FullFlagName(commands.FlagShowContainerLogs), Description: commands.FlagShowContainerLogsUsage},
		{Text: commands.FullFlagName(commands.FlagCRORuntime), Description: commands.FlagCRORuntimeUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPProbe), Description: commands.FlagHTTPProbeUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPProbeCmd), Description: commands.FlagHTTPProbeCmdUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPProbeCmdFile), Description: commands.FlagHTTPProbeCmdFileUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPProbeRetryCount), Description: commands.FlagHTTPProbeRetryCountUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPProbeRetryWait), Description: commands.FlagHTTPProbeRetryWaitUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPProbePorts), Description: commands.FlagHTTPProbePortsUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPProbeFull), Description: commands.FlagHTTPProbeFullUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPProbeExitOnFailure), Description: commands.FlagHTTPProbeExitOnFailureUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPProbeCrawl), Description: commands.FlagHTTPProbeCrawlUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPCrawlMaxDepth), Description: commands.FlagHTTPCrawlMaxDepthUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPCrawlMaxPageCount), Description: commands.FlagHTTPCrawlMaxPageCountUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPCrawlConcurrency), Description: commands.FlagHTTPCrawlConcurrencyUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPMaxConcurrentCrawlers), Description: commands.FlagHTTPMaxConcurrentCrawlersUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPProbeAPISpec), Description: commands.FlagHTTPProbeAPISpecUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPProbeAPISpecFile), Description: commands.FlagHTTPProbeAPISpecFileUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPProbeExec), Description: commands.FlagHTTPProbeExecUsage},
		{Text: commands.FullFlagName(commands.FlagHTTPProbeExecFile), Description: commands.FlagHTTPProbeExecFileUsage},
		{Text: commands.FullFlagName(commands.FlagPublishPort), Description: commands.FlagPublishPortUsage},
		{Text: commands.FullFlagName(commands.FlagPublishExposedPorts), Description: commands.FlagPublishExposedPortsUsage},
		{Text: commands.FullFlagName(commands.FlagKeepPerms), Description: commands.FlagKeepPermsUsage},
		{Text: commands.FullFlagName(commands.FlagRunTargetAsUser), Description: commands.FlagRunTargetAsUserUsage},
		{Text: commands.FullFlagName(commands.FlagCopyMetaArtifacts), Description: commands.FlagCopyMetaArtifactsUsage},
		{Text: commands.FullFlagName(commands.FlagRemoveFileArtifacts), Description: commands.FlagRemoveFileArtifactsUsage},
		{Text: commands.FullFlagName(FlagTag), Description: FlagTagUsage},
		{Text: commands.FullFlagName(FlagImageOverrides), Description: FlagImageOverridesUsage},
		{Text: commands.FullFlagName(commands.FlagEntrypoint), Description: commands.FlagEntrypointUsage},
		{Text: commands.FullFlagName(commands.FlagCmd), Description: commands.FlagCmdUsage},
		{Text: commands.FullFlagName(commands.FlagWorkdir), Description: commands.FlagWorkdirUsage},
		{Text: commands.FullFlagName(commands.FlagEnv), Description: commands.FlagEnvUsage},
		{Text: commands.FullFlagName(commands.FlagLabel), Description: commands.FlagLabelUsage},
		{Text: commands.FullFlagName(commands.FlagVolume), Description: commands.FlagVolumeUsage},
		{Text: commands.FullFlagName(commands.FlagLink), Description: commands.FlagLinkUsage},
		{Text: commands.FullFlagName(commands.FlagEtcHostsMap), Description: commands.FlagEtcHostsMapUsage},
		{Text: commands.FullFlagName(commands.FlagContainerDNS), Description: commands.FlagContainerDNSUsage},
		{Text: commands.FullFlagName(commands.FlagContainerDNSSearch), Description: commands.FlagContainerDNSSearchUsage},
		{Text: commands.FullFlagName(commands.FlagNetwork), Description: commands.FlagNetworkUsage},
		{Text: commands.FullFlagName(commands.FlagHostname), Description: commands.FlagHostnameUsage},
		{Text: commands.FullFlagName(commands.FlagExpose), Description: commands.FlagExposeUsage},
		{Text: commands.FullFlagName(FlagNewEntrypoint), Description: FlagNewEntrypointUsage},
		{Text: commands.FullFlagName(FlagNewCmd), Description: FlagNewCmdUsage},
		{Text: commands.FullFlagName(FlagNewExpose), Description: FlagNewExposeUsage},
		{Text: commands.FullFlagName(FlagNewWorkdir), Description: FlagNewWorkdirUsage},
		{Text: commands.FullFlagName(FlagNewEnv), Description: FlagNewEnvUsage},
		{Text: commands.FullFlagName(FlagNewVolume), Description: FlagNewVolumeUsage},
		{Text: commands.FullFlagName(FlagNewLabel), Description: FlagNewLabelUsage},
		{Text: commands.FullFlagName(FlagRemoveExpose), Description: FlagRemoveExposeUsage},
		{Text: commands.FullFlagName(FlagRemoveEnv), Description: FlagRemoveEnvUsage},
		{Text: commands.FullFlagName(FlagRemoveLabel), Description: FlagRemoveLabelUsage},
		{Text: commands.FullFlagName(FlagRemoveVolume), Description: FlagRemoveVolumeUsage},
		{Text: commands.FullFlagName(commands.FlagExcludeMounts), Description: commands.FlagExcludeMountsUsage},
		{Text: commands.FullFlagName(commands.FlagExcludePattern), Description: commands.FlagExcludePatternUsage},
		{Text: commands.FullFlagName(commands.FlagPathPerms), Description: commands.FlagPathPermsUsage},
		{Text: commands.FullFlagName(commands.FlagPathPermsFile), Description: commands.FlagPathPermsFileUsage},
		{Text: commands.FullFlagName(commands.FlagIncludePath), Description: commands.FlagIncludePathUsage},
		{Text: commands.FullFlagName(commands.FlagIncludePathFile), Description: commands.FlagIncludePathFileUsage},
		{Text: commands.FullFlagName(commands.FlagIncludeBin), Description: commands.FlagIncludeBinUsage},
		{Text: commands.FullFlagName(FlagIncludeBinFile), Description: FlagIncludeBinFileUsage},
		{Text: commands.FullFlagName(commands.FlagIncludeExe), Description: commands.FlagIncludeExeUsage},
		{Text: commands.FullFlagName(FlagIncludeExeFile), Description: FlagIncludeExeFileUsage},
		{Text: commands.FullFlagName(commands.FlagIncludeShell), Description: commands.FlagIncludeShellUsage},
		{Text: commands.FullFlagName(commands.FlagMount), Description: commands.FlagMountUsage},
		{Text: commands.FullFlagName(commands.FlagContinueAfter), Description: commands.FlagContinueAfterUsage},
		{Text: commands.FullFlagName(commands.FlagUseLocalMounts), Description: commands.FlagUseLocalMountsUsage},
		{Text: commands.FullFlagName(commands.FlagUseSensorVolume), Description: commands.FlagUseSensorVolumeUsage},
		{Text: commands.FullFlagName(commands.FlagKeepTmpArtifacts), Description: commands.FlagKeepTmpArtifactsUsage},
		{Text: commands.FullFlagName(FlagBuildFromDockerfile), Description: FlagBuildFromDockerfileUsage},
		{Text: commands.FullFlagName(FlagTagFat), Description: FlagTagFatUsage},
		{Text: commands.FullFlagName(FlagCBOAddHost), Description: FlagCBOAddHostUsage},
		{Text: commands.FullFlagName(FlagCBOBuildArg), Description: FlagCBOBuildArgUsage},
		{Text: commands.FullFlagName(FlagCBOLabel), Description: FlagCBOLabelUsage},
		{Text: commands.FullFlagName(FlagCBOTarget), Description: FlagCBOTargetUsage},
		{Text: commands.FullFlagName(FlagCBONetwork), Description: FlagCBONetworkUsage},
		{Text: commands.FullFlagName(FlagCBOCacheFrom), Description: FlagCBOCacheFromUsage},
	},
	Values: map[string]commands.CompleteValue{
		//NOTE: with FlagPull target complete needs to check remote registries too
		commands.FullFlagName(commands.FlagPull):                   commands.CompleteBool,
		commands.FullFlagName(commands.FlagShowPullLogs):           commands.CompleteBool,
		commands.FullFlagName(commands.FlagTarget):                 commands.CompleteTarget,
		commands.FullFlagName(FlagShowBuildLogs):                   commands.CompleteBool,
		commands.FullFlagName(commands.FlagShowContainerLogs):      commands.CompleteBool,
		commands.FullFlagName(commands.FlagPublishExposedPorts):    commands.CompleteBool,
		commands.FullFlagName(commands.FlagHTTPProbe):              commands.CompleteTBool,
		commands.FullFlagName(commands.FlagHTTPProbeCmdFile):       commands.CompleteFile,
		commands.FullFlagName(commands.FlagHTTPProbeExecFile):      commands.CompleteFile,
		commands.FullFlagName(commands.FlagHTTPProbeFull):          commands.CompleteBool,
		commands.FullFlagName(commands.FlagHTTPProbeExitOnFailure): commands.CompleteBool,
		commands.FullFlagName(commands.FlagHTTPProbeCrawl):         commands.CompleteTBool,
		commands.FullFlagName(commands.FlagHTTPProbeAPISpecFile):   commands.CompleteFile,
		commands.FullFlagName(commands.FlagKeepPerms):              commands.CompleteTBool,
		commands.FullFlagName(commands.FlagRunTargetAsUser):        commands.CompleteTBool,
		commands.FullFlagName(commands.FlagRemoveFileArtifacts):    commands.CompleteBool,
		commands.FullFlagName(commands.FlagNetwork):                commands.CompleteNetwork,
		commands.FullFlagName(commands.FlagExcludeMounts):          commands.CompleteTBool,
		commands.FullFlagName(commands.FlagPathPermsFile):          commands.CompleteFile,
		commands.FullFlagName(commands.FlagIncludePathFile):        commands.CompleteFile,
		commands.FullFlagName(FlagIncludeBinFile):                  commands.CompleteFile,
		commands.FullFlagName(FlagIncludeExeFile):                  commands.CompleteFile,
		commands.FullFlagName(commands.FlagIncludeShell):           commands.CompleteBool,
		commands.FullFlagName(commands.FlagContinueAfter):          commands.CompleteContinueAfter,
		commands.FullFlagName(commands.FlagUseLocalMounts):         commands.CompleteBool,
		commands.FullFlagName(commands.FlagUseSensorVolume):        commands.CompleteVolume,
		commands.FullFlagName(commands.FlagKeepTmpArtifacts):       commands.CompleteBool,
	},
}
