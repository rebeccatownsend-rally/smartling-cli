package main

import (
	"fmt"
)

const formatOptionHelp = `
This command supports advanced formatting via --format flag with full
support of Golang templates (https://golang.org/pkg/text/template).

Special formatting functions are available:

  > {{name <variable>}} — return file URI without extension for specified
    <variable>;
  > {{ext <variable}} — return extension from file URI for specified <variable>;
`

const authenticationOptionsHelp = `
  --user <user>
    Specify user ID for authentication.

  --secret <secret>
    Specify secret token for authentication.

  -a --account <account>
    Specify account ID.
`

const globPatternHelp = `argument support globbing with following patterns:

  > ** — matches any number of any chars;
  > *  — matches any number of chars except '/';
  > ?  — matches any single char except '/';
  > [xyz]   — matches 'x', 'y' or 'z' charachers;
  > [!xyz]  — matches not 'x', 'y' or 'z' charachers;
  > {a,b,c} — matches alternatives a, b or c;`

const initHelp = `smartling-cli init — create config file interactively.

Walk down common config file parameters and fill them through dialog.

Init process will inspect if config file already exists and if it is, it will
be loaded as default values, so init can be used sequentially without config
is lost.

Options like --user, --secret, --account and --project can be used to specify
config values prior dialog:

  smartling-cli init --user=your_user_id

Also, --dry-run option can be used to just look at resulting config without
overwritting anything:

  smartling-cli init --dry-run

By default, smartling.yml file in the local directory will be used as target
config file, but it can be overriden by using --config option:

  smartling-cli init --config=/path/to/project/smartling.yml


Available options:
  -c --config <file>
    Specify config file to operate on. Default: smartling.yml

  --dry-run
    Do not overwrite config file, only output to stdout.

Default config values can be passed via following options:` +
	authenticationOptionsHelp + `
  -p --project <project>
    Specify default project.
`

const projectsListHelp = `smartling-cli projects list — list projects from account.

Command will list projects from specified account in tabular format with
following information:

  > Project ID
  > Project Description
  > Project Source Locale ID

Only project IDs will be listed if --short option is specified.

Note, that you should specify account ID either in config file or via --account
option to be able to see projects list.


Available options:
  -s --short
    List only project IDs.
` + authenticationOptionsHelp

const projectsInfoHelp = `smartling-cli projects info — show detailed project info.

Displays detailed information for specific project.

Project should be specified either in config or via --project option.


Available options:` + authenticationOptionsHelp

const projectsLocalesHelp = `smartling-cli projects locales — list target locales.

Lists target locales from specified project.

To list only locale IDs --short option can be used.
` + formatOptionHelp + `
Following variables are available:

  > .LocaleID — target locale ID to translate into;
  > .Description — human-readable locale description;
  > .Enabled — true/false specifying is locale active or not;


Available options:
  -p --project <project>
    Specify project to use.

  -s --short
    List only locale IDs.

  --format
    Use specific output format instead of default.
` + authenticationOptionsHelp

const filesListHelp = `smartling-cli files list — list files from project.

Lists all files from project or only files which matches specified uri.

Note, that by default listing is limited to 500 items in Smartling API,
so several requests may be needed to obtain full file list, which will
take some time.

List command will output following fields in tabular format by default:

  > File URI;
  > Last uploaded date;
  > File Type;
` + formatOptionHelp + `
Following variables are available:

  > .FileURI — full file URI in Smartling system;
  > .FileType — internal Smartling file type;
  > .LastUploaded — timestamp when file was last uploaded;
  > .HasInstructions — true/false if file has translation instructions;

<uri> ` + globPatternHelp + `


Available options:
  -p --project <project>
    Specify project to use.

  -s --short
    List only file URIs.

  --format <format>
    Override default listing format.
` + authenticationOptionsHelp

const filesPullHelp = `smartling-cli files pull — downloads translated files from project.

Downloads files from specified project into local directory.

It's possible to download only specific files by file mask, to download source
files with translations, to download file to specific directory or to download
specific locales only.

If special value of "-" is specified as <uri>, then program will expect
to read files list from stdin:

  cat files.txt | smartling-cli files pull -

<uri> ` + globPatternHelp + `

If --locale flag is not specified, all available locales are downloaded. To
see available locales, use "status" command.

To download files into subdirectory, use --directory option and specify
directory name you want to download into.

To download source file as well as translated files specify --source option.

Files will be downloaded and stored under names used while upload (e.g. File
URI). While downloading translated file suffix "_<locale>" will be appended to
file name before extension. To override file format name, use --format option.
` + formatOptionHelp + `
Following variables are available:

  > .FileURI — full file URI in Smartling system;
  > .Locale — locale ID for translated file and empty for source file;


Available options:
  -p --project <project>
    Specify project to use.

  --source
    Download source files along with translated files.

  —d ——directory <dir>
    Download files into specified directory.

  --format <format>
    Specify format for download file nmae.

  --progress <percents>
    Specify minimum of translation progress in percents.
	By default that filter does not apply.

  --retrieve <type>
    Retrieval type according to API specs:
    > pending — returns any translations, including non-published ones);
    > published — returns only published translations;
    > pseudo — returns modified version of original text with certain
               characters transformed;
    > contextMatchingInstrumented — to use with Chrome Context Capture;
` + authenticationOptionsHelp

const filesPushHelp = `smartling-cli files push <file> [<uri>] [--type <type>] [--branch (@auto|<branch name>)] [--authorize|--locale <locale>] [--directory <work dir>] [--directive <smartling directive>]

Uploads files designated for translation.

One or several files can be pushed.

When pushing single file, <uri> can be specified to override local path.
When pushing multiple files, they will be uploaded using local path as URI.
If no file specified in command line, config file will be used to lookup
for file masks to push.

To authorize all locales, use --authorize option.

To authorize only specific locales, use one or more --locale.

To prepend prefix to all target URIs, use --branch option. Special
value "@auto" can be used to tell that tool should try to took current git
branch name as value for --branch option.

File type will be deduced from file extension. If file extension is unknown,
type should be specified manually by using --type option. That option also
can be used to override detected file type.

<file> ` + globPatternHelp + `


Available options:
  -p --project <project>
    Specify project to use.

  --authorize
    Authorize all available locales. Incompatible with --locale option.

  --locale <locale>
    Authorize speicified locale only. Can be specified several times.
    Incompatible with --authorize option.

  --branch <branch>
    Prepend specified prefix to target file URI.

  --type <type>
    Override automatically detected file type.
` + authenticationOptionsHelp

const filesStatusHelp = `smartling-cli files status — show files status from project.

Lists all files from project along with their translation progress into
different locales.

Status command will check, if files are missing locally or not.

Command will list projects from specified account in tabular format with
following information:

  > File URI
  > File Locale
  > File Status on Local System
  > Translation Progress
  > Strings Count
  > Words Count

If no <uri> is specified, all files will be listed.

To list files status from specific directory, --directory option can be used.

To override default file name format --format can be used.
` + formatOptionHelp + `
Following variables are available:

  > .FileURI — full file URI in Smartling system;
  > .Locale — locale ID for translated file and empty for source file;

<uri> ` + globPatternHelp + `


Available options:
  -p --project <project>
    Specify project to use.

  --directory <directory>
    Check files in specific directory instead of local directory.

  --format <format>
    Specify format for listing file names.
` + authenticationOptionsHelp

const filesDeleteHelp = `smartling-cli files delete — removes files from project.

Removes files from project according to specified pattern.

<uri> ` + globPatternHelp + `

If special value of "-" is specified as <uri>, then program will expect
to read files list from stdin:

  cat files.txt | smartling-cli files delete -

Available options:
  -p --project <project>
    Specify project to use.
` + authenticationOptionsHelp

const filesRenameHelp = `smartling-cli files rename — rename specified file.

Renames specified file URI into new file URI.

Available options:
  -p --project <project>
    Specify project to use.
` + authenticationOptionsHelp

const importHelp = `smartling-cli import — import file translations.

Import pre-existent file translations into Smartling. Note, that
original file should be pushed prior file translations are imported.

Either --published or --post-translation should present to specify state
of imported translation.  Value indicates the workflow state to import the
translations into. Content will be imported into the language's default
workflow.

--overwrite option can be used to replace existent translations.

Available options:
  --published
    The translated content is published.

  --post-translation
   The translated content is imported into the first step after translation
   If there are none, it will be published.

  --overwrite
    Overwrite existing translations.
` + authenticationOptionsHelp

func showHelp(args map[string]interface{}) {
	switch {
	case args["init"].(bool):
		fmt.Print(initHelp)

	case args["projects"].(bool):
		switch {
		case args["list"].(bool):
			fmt.Print(projectsListHelp)
		case args["info"].(bool):
			fmt.Print(projectsInfoHelp)
		case args["locales"].(bool):
			fmt.Print(projectsLocalesHelp)
		}

	case args["files"].(bool):
		switch {
		case args["list"].(bool):
			fmt.Print(filesListHelp)
		case args["pull"].(bool), args["get"].(bool):
			fmt.Print(filesPullHelp)
		case args["push"].(bool):
			fmt.Print(filesPushHelp)
		case args["status"].(bool):
			fmt.Print(filesStatusHelp)
		case args["delete"].(bool):
			fmt.Print(filesDeleteHelp)
		case args["rename"].(bool):
			fmt.Print(filesRenameHelp)
		case args["import"].(bool):
			fmt.Print(importHelp)
		}

	default:
		fmt.Print(usage)
	}
}
