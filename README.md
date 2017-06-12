<p>It's Util of wechat API implments.<br />
This util doesn't contain HTTP server, only API for generate URL or handle response
</p>

<p>
  Contains API implements: <br />
  1. token auth   <br />
  2. user token auth <br />
  3. js token auth <br />
  4. payment API <br />
</p>
<p>
Example:
Initalize glwapi configuration:
<code>

    g := glwapi.D() 

    ret := g.Init(data.Appid, data.Secret, data.MchId, data.MchApiKey)
    if ret != true {
         log.Fatal(" initalize glwapi failed")
    }
</code>

Get Token:

<code>

 g := glwapi.D()

 if g.IsInit() == false {

       LE("glwapi doesn't inital yet")

       return
 }

 token_url := g.TokenURL()

 r, err :=  http.Get(token_url)

 if err == nil {

       defer r.Body.Close()

       rdata, _ := ioutil.ReadAll(r.Body)

       _, e := g.HandleTokenAuthResp([]byte(rdata))

       if e != nil {

            LE("Handle auth response error:%s", e)

       } else {

            LI("Get auth token successfully: %s\n", g)
       }

 }

</code>

	
</p>
