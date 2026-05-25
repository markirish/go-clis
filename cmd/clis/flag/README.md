# Go CLI Comparison

This repository seeks to evaluate different first- and third-party options for creating command-line interfaces (CLIs) with Go. 

## Evaluation Criteria

### Functionality

In each sub-directory, a CLI is created with the same interface spec (based on and reduced from another repo of mine, oc-inspector). This spec requires the CLI to have the following commands (verbs) on the following object types (nouns):

* list
  * pods
  * deployments
  * services
  * nodes
  * events
* describe
  * pods
  * deployments
  * services
  * nodes
* diff
  * pods
  * deployments
  * services

The CLI library must be able to use these verb-noun contexts (VNCs) to enforce option tpying, use default values for applicable options, and determine which passed options are valid (or invalid), required, excluded, dependent on one-another, and which combinations are mutually exclusive.

The following options should be evaluated correctly:

#### Global Options

These options should always be valid

`--verbose, -v`

`bool`\
Turning on verbose output to provide more fine-grained tracing

`--namespace, -n`

`string`\
Limit interaction with the cluster to just the passed namespace

`--timeout, --request-timeout`

`time.Duration`\
The timeout for any given request to the cluster

`--global-timeout`

`time.Duration`\
The timeout for all of the work the program is requested to do

`--output, -o`

`enum: (table, json, yaml)`\
Structures the output so it can be more easily consumed by the caller

#### Verb-Level Options

These options should be available to all calls to the verb they are listed under, no matter which noun is called

#### `list` Verb-Level Options

These options should be valid whenever the `list` verb is used

`--filter`

`--sort-by`

`--watch, -w`

`--wide`

#### `list` Noun-Level Options

These options are only valid only for the noun they are listed under when calling the `list` verb

##### `list pods`

`--phase`

`--node`

`--restarts-gt`

##### `list deployments`

`--unavailable`

`--min-replicas`

##### `list services`

`--type`

##### `list nodes`

`--node`

`--not-ready`

##### `list events`

`--type`

`--since`

---

#### `describe` Verb-Level Options

`--filter`

`--name`

`--show-events`

#### `describe` Noun-Level Options

##### `describe pods`

`--show-containers`

`--show-logs`

##### `describe deployments`

`--show-replicas`

`--show-pods`

##### `describe services`

`--show-endpoints`

`--show-selectors`

##### `describe nodes`

`--show-pods`

`--show-capacity`

---

#### `diff` Verb-Level Options

`--file, -f`

`--name`

`--ignore-metadata`

#### `diff` Noun-Level Options

##### `diff pods`

`--ignore-image-tag`

`--container`

##### `diff deployments`

`--ignore-replicas`

`--ignore-strategy`

##### `diff services`

`--ignore-cluster-ip`

`--ignore-node-port`

---



### Usability

In addition to actually being able to implement the above interface requirements, it 




## Terminology

**verb**: The first positional argument after the name of the CLI binary. Indicates what the caller is trying to do.

e.g.:

```bash
$ dnf install golang
      ^^^^^^^
      verb
```

```bash
$ ocm get pods
      ^^^
      verb
```

**noun**: The second positional argument, found immediately after the verb. Indicates the type of domain object that the caller is attepting to *verb*.

e.g.:

```bash
$ dnf install golang
              ^^^^^^
              noun
```

```bash
$ ocm get pods
          ^^^
          noun
```

**VNC** - The **verb/noun context**. This context dictates which command options are valid or invalid. Some combinations of verb + noun are 

**caller** - The entity that invoked the CLI program. This may be a user manually entering commands into a terminal, or a computer 

### Required Options

The ability of the options parser to be aware of required options for a given VNC.

### Dependent Options

If one option is defined, another option becomes valid or required

### VNC Options

The ability for the options parser to determine if options are valid for the given verb/noun context.

### Mutually Exclusive Options

The ability for the options parser to determine if two or more passed options conflict with one another.

