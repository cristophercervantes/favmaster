
from setuptools import setup, find_packages

with open("README.md", "r", encoding="utf-8") as fh:
    long_description = fh.read()

setup(
    name="favmaster",
    version="1.0.0",
    author="Cristopher",
    author_email="cristopher@example.com",
    description="A powerful tool for extracting MMH3, SHA256, and MD5 hashes from images and favicons",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/cristophercervantes/favmaster",
    packages=find_packages(),
    classifiers=[
        "Development Status :: 4 - Beta",
        "Intended Audience :: Developers",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
        "Programming Language :: Python :: 3",
    ],
    python_requires=">=3.7",
    install_requires=[
        "requests>=2.25.1",
        "beautifulsoup4>=4.9.3",
        "mmh3>=3.0.0",
    ],
    entry_points={
        "console_scripts": [
            "favmaster=favmaster.cli:main",
        ],
    },
)

