# `proto/`

This directory contains the protocol buffer definitions for all NeCode components. 
We follow the convention of having a single `proto` folder to make cross-component imports and codegen easier.

We use [Buf](https://buf.build) to lint and generate protocol buffers.

[style guide](https://docs.buf.build/best-practices/style-guide).

## Defining APIs

We define APIs as gRPC services. All APIs should be defined centrally in this directory, but implemented in their respective sub-package of the monorepo. 

## Generating

After changing a `.proto` file, you should re-generate the bindings:

1. Install Buf if you haven't already ([docs](https://docs.buf.build/installation))
2. From the repo root, run:
```bash
make proto.generate
```