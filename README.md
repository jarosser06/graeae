Graeae
======
Quick and easy way to share files by uploading them to S3 and creating
a presigned URL. There is an expectation that you will have access to the
S3 Bucket you are trying to use.

For a basic bucket configuration to use there is a CloudFormation template with a Lifecycle policy that expires
objects after 7 days.

Example:
```shell
graeae -bucket misc.j4r.wtf ./foo.zip
```

Example setting the amount of time the pre-signed URL is active.
```shell
graeae -bucket misc.j4r.wtf -valid 20 ./foo.zip
```
