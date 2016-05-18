

# rdm3u8
 remove domain in m3u8 and upload to bucket

## 用法：
```
$./rdm3u8 <config_file_path>
```

## 配置文件：
```
{
    "domain": <m3u8.domain.com>,
    "access_key": <access_key>,
    "secret_key": <secret_key>,
    "bucket": <m3u8_bucket>,
    "bucket_bak": <m3u8_bak_bucket>,
    "keys_file_loc": <m3u8_names_file_addr>
}
```

domain: 需要处理的m3u8文件所在的域名。

access_key: m3u8文件所在账号的 access_key。

secret_key: m3u8文件所在账号的 secret_key。

bucket: m3u8文件所在的空间。

bucket_bak: 可选， 设置后，在处理m3u8之前会将 m3u8 文件复制到设置的 bucket_bak 中

keys_file_loc: 需要处理的 m3u8 文件的 文件名， 每行一个。





