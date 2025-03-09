// firefox config for neko
lockPref("browser.tabs.closeWindowWithLastTab", false);
lockPref("app.update.auto", false);
lockPref("app.update.enabled", false);
lockPref("app.update.silent", true);
lockPref("browser.cache.disk.capacity", 1000);
lockPref("browser.download.useDownloadDir", false);
lockPref("browser.rights.3.shown", true);
lockPref("browser.search.update", false);
lockPref("browser.shell.checkDefaultBrowser", false);
lockPref("extensions.update.enabled", false);
lockPref("plugin.default_plugin_disabled", false);
lockPref("plugin.scan.plid.all", true);
lockPref("plugins.hide_infobar_for_missing_plugin", true);
lockPref("profile.allow_automigration", false);
lockPref("signon.prefillForms", false);
lockPref("signon.rememberSignons", false);
lockPref("xpinstall.enabled", false);
lockPref("xpinstall.whitelist.required", true);
lockPref("browser.download.manager.retention", 0);
lockPref("browser.download.folderList",	2);
lockPref("browser.download.forbid_open_with", true);
lockPref("browser.safebrowsing.downloads.enabled", false);
lockPref("browser.safebrowsing.downloads.remote.enabled",	false);
lockPref("browser.helperApps.alwaysAsk.force",	false);
lockPref("browser.helperApps.neverAsk.saveToDisk",	"application/zip,application/octet-stream,image/jpeg,application/vnd.ms-outlook,text/html,application/pdf");
lockPref("browser.helperApps.neverAsk.openFile",	"application/zip,application/octet-stream,image/jpeg,application/vnd.ms-outlook,text/html,application/pdf");
//lockPref("browser.newtabpage.activity-stream.default.sites",	"");
// dark mode
lockPref("reader.color_scheme", "dark");
lockPref("devtools.theme", "dark");
lockPref("ui.systemUsesDarkTheme", 1);
lockPref("lightweightThemes.usedThemes","[]");
lockPref("lightweightThemes.selectedThemeID", "firefox-compact-dark@mozilla.org");
lockPref("extensions.activeThemeID", "firefox-compact-dark@mozilla.org");
lockPref("browser.theme.toolbar-theme", 0);
lockPref("browser.in-content.dark-mode", true);

// mousewheel fix
lockPref("mousewheel.default.delta_multiplier_x", 45);
lockPref("mousewheel.default.delta_multiplier_y", 45);
lockPref("mousewheel.default.delta_multiplier_z", 45);

// disable vie wbtn
lockPref("browser.tabs.firefox-view", false);
lockPref("browser.tabs.firefox-view-next", false);
lockPref("browser.tabs.tabmanager.enabled", false);

// telemetry
lockPref("devtools.onboarding.telemetry.logged", false);
lockPref("toolkit.telemetry.updatePing.enabled", false);
lockPref("browser.newtabpage.activity-stream.feeds.telemetry", false);
lockPref("browser.newtabpage.activity-stream.telemetry", false);
lockPref("browser.newtabpage.activity-stream.feeds.discoverystreamfeed", false);
lockPref("browser.newtabpage.activity-stream.showSponsored", false);
lockPref("browser.ping-centre.telemetry", false);
lockPref("toolkit.telemetry.bhrPing.enabled", false);
lockPref("toolkit.telemetry.enabled", false);
lockPref("toolkit.telemetry.firstShutdownPing.enabled", false);
lockPref("toolkit.telemetry.hybridContent.enabled", false);
lockPref("toolkit.telemetry.newProfilePing.enabled", false);
lockPref("toolkit.telemetry.reportingpolicy.firstRun", false);
lockPref("toolkit.telemetry.shutdownPingSender.enabled", false);
lockPref("toolkit.telemetry.unified", false);
lockPref("toolkit.telemetry.archive.enabled", false);
lockPref("datareporting.healthreport.uploadEnabled", false);
lockPref("datareporting.policy.dataSubmissionEnabled", false);
lockPref("datareporting.sessions.current.clean", false);

// disable activity stream
lockPref("browser.newtabpage.enabled", false); 
lockPref("browser.newtab.preload", false); 
lockPref("browser.newtabpage.activity-stream.aboutHome.enabled", false);

// disable geo
lockPref("geo.provider.ms-windows-location", false);
lockPref("geo.provider.use_gpsd", false);
lockPref("geo.provider.use_corelocation", false);
lockPref("geo.provider.use_geoclue", false);
lockPref("browser.region.network.url", "");
lockPref("browser.region.update.enabled", false);

