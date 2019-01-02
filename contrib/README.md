# Contribution Guide

## Summary

This project follows [Semantic Versioning 2.0](https://semver.org/)

## Publishing a version

1. Update the version field in `.evg.yml`
2. Commit your changes, push upstream, and wait for the build to pass

  ```bash
  git commit -m "Bump version to 1.x.x"
  git push upstream HEAD
  ```

3. *After* a successful build, run

  ```bash
  ./contrib/bump_version.sh 1.x.x
  ```

  **NOTE** this assumes that you have the `aws` CLI installed

4. Update the `CURRENT` release file in S3 so that the correct version of the CLI can be downloaded via the Stitch Admin UI

  First, download the [CURRENT](https://s3.console.aws.amazon.com/s3/object/stitch-clis/versions/cloud-prod/CURRENT?region=us-east-1&tab=overview) file and rename to `CURRENT` (**NOT** `CURRENT.json`)

  Then, replace the body of the `CURRENT` file with the output printed after running `./contrib/bump_version.sh 1.x.x` (it should be copied to your clipboard if you're on OS X)

  Finally, upload the updated `CURRENT` file to the [stitch-clis/versions/cloud-prod](https://s3.console.aws.amazon.com/s3/buckets/stitch-clis/versions/cloud-prod/?region=us-east-1&tab=overview) bucket

  ### IMPORTANT

  Ensure that when uploading the `CURRENT` file that you:

    * Select `Grant public read access to this object(s)` under `Manage public permissions`
    * Set the `Content-Type` header to `application/json` under `Metadata`

5. Push your changes upstream with `git push upstream --follow-tags`
6. Run `npm publish` to publish
