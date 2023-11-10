import os
import sys


project_name = "URLShortner"
build_directory = "./build"
src_path = "./cmd/URLShortner/main.go"

project_build = os.path.join(build_directory, project_name)

os.system(f"go build -o {project_build} {src_path}")