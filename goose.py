import argparse
import os

def parse_env() -> dict:
    d = {}
    with open(".env", 'r', encoding="utf-8") as env:
        lines = env.readlines()
        for line in lines:
            line = line.split('=')
            d[line[0]] = line[1].strip()
    return d


migrations_dir = "./internal/storage/postgresql/migrations/schema"
up_migration = "up"
down_migration = "down"

def parse_flags() -> dict:
    parser = argparse.ArgumentParser()
    parser.add_argument('--up', dest = 'up', action = 'store_true', help = 'up migrations')
    parser.add_argument('--down', dest = 'down', action = 'store_true', help = 'down migrations')
    flags = parser.parse_args().__dict__
    return flags

def migration_type(flags: dict) -> str:
    if flags[up_migration]:
        return up_migration
    return down_migration

if __name__ == "__main__":

    env = parse_env()

    flags = parse_flags()

    if flags["up"] and flags["down"] or not flags["up"] and not flags["down"]:
        print("--up of --down required.")
    else:
        migration = migration_type(flags)
        psql = env["DB_URL_GOOSE"]
        os.system(f"goose -dir {migrations_dir} postgres {psql} {migration}")
