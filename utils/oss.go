package utils
//import "github.com/astaxie/beego"
//
//var ossHost, ossId, ossKey, ossBucket string
//var ossInternal bool
//var c *oss.Client
//
//func ossInit(){
//	ossInternal  = beego.AppConfig.String("oss_internal")
//	if ossInternal {
//		ossHost = beego.AppConfig.String("oss_put_host_internal")
//	}else{
//		ossHost = beego.AppConfig.String("oss_put_host")
//	}
//	ossId = beego.AppConfig.String("oss_id")
//	ossKey = beego.AppConfig.String("oss_key")
//	ossBucket = beego.AppConfig.String("oss_bucket")
//
//	c = oss.NewClient(ossHost, ossId, ossKey, 10)
//}
//
//func OssStore(opath, fpath string)error{
//	ossInit()
//	err := c.PutObject(ossBucket + "/" + opath, fpath)
//	return err
//}
//
//func OssDelete(opath string)error{
//	ossInit()
//	err := c.DeleteObject(ossBucket + "/" + opath)
//	return err
//}
//
//func OssGetUrl(opath string)string{
//	return beego.AppConfig.String("ooss_get_host") + "/" + opath
//}