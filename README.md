# ThreatDefender

Note: This is still in alpha release, this might break over time.

## About
<img src="assets/img.png" width="200" height="200">


ThreatDefender is a tool to scrape for potential dangerous threats faced on the internet. The list of threats scraped at
the moment is:

- Emails
    - Disposable
    - Generic
- IPs
    - VPN
    - Spam
    - Proxy
    - Tor

This tool saves those threats on multiple different databases and uses REST API to outdata

# Usage

## Migrate

**BEFORE YOU RUN THIS**, You need to migrate the database:

`make migrate`

## How to run

To run it on your local computer:

`make run`

### Changing ports and configs

All configurations are in the config files in the directory "config". You can add your own environment variables here.

## Apis

REST API & gRPC is enabled.

Take a look at the documentation here: 

## Work in progress

Lots of features are being worked on.

## Roadmap

I would like a discussion going on the potential expansion of the tool.

I would like this tool to detect all modern threats.