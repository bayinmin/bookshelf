#!/usr/bin/python

import requests
import argparse
from colorama import Fore, Style

def debug_mode_is_on(url):
    try:
        console_url = url + "/console"
        print("Checking debug console is accessible via: " + console_url)
        r = requests.get(console_url)
        if(r.status_code == 200):
            print(Fore.RED + "** debug console path is accessbile. You should try RCE exploit at " + console_url)
        else:
            print(Fore.GREEN + console_url + " is not accessible. status code: " + str(r.status_code))
    except Exception as err:
        print(f"Unexpected {err=}, {type(err)=}")
    print(Style.RESET_ALL)
    print("DONE")

if __name__=='__main__':
    p = argparse.ArgumentParser()
    p.add_argument('--host', type=str, required=True, help='e.g. http://127.0.0.1/')
    args = p.parse_args()

    debug_mode_is_on(args.host)
