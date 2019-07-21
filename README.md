# Remote live config

Software may often look like this

    main()
    settings = load (json-config-file)

Utilities like [viper][](golang), [anyconfig][](python), or [node-config][](nodejs) allow you to read and compose config from multiple locations and environment variables.

There has long been a [consideration](https://softwareengineering.stackexchange.com/questions/179572/should-i-use-a-config-file-or-database-for-storing-business-rule) of whether to store certain application settings in config files, or directly in a database.

This demonstration (repo) blends the two, by using a small utility to sync data in a database, to a config file on disk. This decouples the two approaches, so that the code does not have to know about the database connection.

[ program ] ---  [ file-config ] --- [ sync-tool ] --- [ database ]

The syncronization is made more efficient at two points:

 - The database used is Firestore, which has a live data watch pattern, which pushes data to listening clients only on data change, no polling.
 - System-level filesystem events can be used to monitor files for changes on disk (again, no polling).

## Try it out



### Permissions

The current design makes the assumption that this pattern may be used by several services at the same time, and you don't want to allow services to modify config. You need to set up the service account you use to have read-only permission to datastore. If you are using Datastore/Firestore as a primary data-source, it should live in a different project.

You can use a json-key file and set GOOGLE_APPLICATION_CREDENTIALS env var to the path, or use compute implicit service accounts in services like VMs, Cloud Run, etc.

If the sync tool determines it has write permissions, it will quit.

## Extensions

### Extending read-once libraries with live reloading

[viper][] allows hot-reloading of files if they change on disk while the referenced python and nodejs projects do not. Included are relatively simple examples of how to hot-reload the json settings. This can be applied as a live-patching of settings loaded through these other utilities.

### Extending for multiple files

Instead of the utility syncing just one file, it would be relatively easy to watch the whole collection and sync either all files, or any file that already exists on disk.

### Show in Containers/Kubernetes

Show how to use with [multi-process container](https://github.com/ahmetb/multi-process-container).

Update a kubernetes configmap, or as alternative to config maps, but with the sync process running as sidecar container in a pod.

## Compared to Git?

 - There are configurations you want to track in git, especially as it relates to infrastructure configuration, as this creates a strong version history.
 - Git is less suited for more rapidly changing values that are app-specific.
 - Git is ideal whenever there are settings more directly tied to a generation of code that makese sense to always store together. But if they are extremely coupled, then maybe this shouldn't be in a config file, but in the code if it is at the app layer.

## Compared Environment Variables?

 - Environment variables are only read on process start, and generally are not, or can not be reloaded.
 - Env vars are generally considered a sound way to expose secrets to a process (though how they get set is subject to a multitude of strategies). While this system, as built, has limited security.
 - Env variables have to be changed/updated per environment, where these settings can be updated quickly across a range of environments.

## Compared to a database?

This is for config and settings, which means a pretty finite and reasonably bounded set. This will depend on your application.  Is the conversion rate of 50 different currencies data or settings?  There is some eye of the beholder work here. As stated earlier - this blends the long running debate about storing settings in files or a database, allowing the application to use a file-based interface, to data remotely managed in a database. Clearly this should not be used for primary application data.

[viper]: https://github.com/spf13/viper

[anyconfig]: https://github.com/lorenwest/node-config/wiki/Configuration-Files

[node-config]: https://github.com/ssato/python-anyconfig
