<!DOCTYPE html>
<html lang="zh-cn">
  <head>
    <title>Teletraan file upload demo</title>
  </head>
  <body>
    <form
      enctype="multipart/form-data"
      action="http://localhost:4001/ocr"
      method="post">

      <label>上传图片</label>
      <input type="file" name="imgFile">
      <br><label>填写userId, 如果你日后想要查找记录的话, 可以用userId来查找</label>
      <input type="text" name="userId">
      <br><label>你是否同意teletraan将你上传的文件匿名保存(ksuid标识), 用来提高服务质量</label>
  	<label for="isAgree">Yes</label>
  	<input type="radio" name="isAgree" value="yes" checked="checked">
  	<label for="isAgree">No</label>
  	<input type="radio" name="isAgree" value="no">
    <br><label>选择需要识别的语言, 目前系统默认支持(简体中文 chi_sim, 英文 eng, 日文 jpn), 如果需要多个一起识别, 请用';', 分隔,不要有空格. 例如: 一张图片中同时包含简体中文和日文 chi_sim;jpn</label>
    <input type="text" name="wantRecognizeLans" value="chi_sim">
    <br><input type="submit" value="开始上传">
    </form>
    <form action="http://localhost:4001/query" method="get">
    <label>查询历史记录, 请输入你需要查询的UserId</label>
    <input type="text" name="userId">
    <input type="submit" value="开始查询">
    </form>
  </body>
</html>
