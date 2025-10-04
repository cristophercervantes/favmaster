from setuptools import setup, find_packages

setup(
    name="favmaster",
    version="0.1.1",
    packages=find_packages(),
    install_requires=[
        "requests>=2.28.0",
        "beautifulsoup4>=4.11.0",
        "mmh3>=4.0.0",
    ],
    entry_points={
        "console_scripts": [
            "favmaster = favmaster.main:main",
        ],
    },
    author="Cristopher",
    description="Compute MMH3 (Shodan/FOFA/Censys/CriminalIP-compatible), MD5, and SHA256 hashes from a favicon",
    url="https://github.com/cristophercervantes/favmaster",  
    classifiers=[
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
    ],
)
