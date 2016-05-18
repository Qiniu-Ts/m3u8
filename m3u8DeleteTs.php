<?php
	require_once './autoload.php';
	require_once'./config.php';
	use Qiniu\Auth;
	use Qiniu\Storage\BucketManager;
	$auth = new Auth(Config::ACCESS_KEY, Config::SECRET_KEY);
	$bucketMgr = new BucketManager($auth);
	$bucket = 'pili';
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
						$value=substr($value, 1);
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
	     else
		     {
						$pos = strpos($newstring, '/',7);
						$value=substr($newstring, $pos+1);
						$err = $bucketMgr->delete($bucket, $value);
				if ($err !== null) {
				    	array_push($at, $value."=>faill");
				} else {
						array_push($at, $value."=>success");
				}
		     }

			}
	}

	// huxicong
	echo "<pre>";
	var_dump($at);
	echo "</pre>";
	file_put_contents("./de.txt", json_encode($da));
?>
