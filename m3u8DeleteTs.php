<?php
/*
目前该api是配合PHP sdk使用的
列出M3u8列表的TS文件
使用BucketManager->delete();来循环删除
*/
	require_once './autoload.php';
	require_once'./config.php';
	use Qiniu\Auth;
	use Qiniu\Storage\BucketManager;
	$auth = new Auth($ACCESS_KEY, $SECRET_KEY)//这里传入是你的ak,sk
	$bucketMgr = new BucketManager($auth);
	$bucket = 'bucket';//这里是你的空间名
	$data=file_get_contents("http://pili-static.pili.echohu.top/recordings/z1.1314xicong.56cc6c21eb6f9275bb0149ff/kkkkk.m3u8");//获取M3u8的文件列表
	$da=explode(PHP_EOL, $data);
	$at=array();
	$dd=true;
		if (preg_match("/^(http:\/\/)?([^\/]+)/i",$da[0]))
		{
			$dd=false;
		}

	foreach ($da as $key => $value) 
	{
		if (preg_match("/\.(?:csv|ts)$/i", $value))
		{
		if ($dd)
			{

				//没有域名前缀的TS文件的删除
						$value=substr($value, 1);
		     }
	     else
		     {
		     	//有带域名前缀的TS文件的删除
						$pos = strpos($newstring, '/',7);
						$value=substr($newstring, $pos+1);
		     }
		     	$err = $bucketMgr->delete($bucket, $value);
				if ($err !== null) 
				{
				    	array_push($at, $value."=>faill");
				}
				 else 
				{
						array_push($at, $value."=>success");
				}

			}
	}

	// huxicong
	echo "<pre>";
	var_dump($at);
	echo "</pre>";
	file_put_contents("./de.txt", json_encode($da));
?>
