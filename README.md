# go_helper

是一个对go 的 gorm 升级版的插件 将where条件直接绑定 struct 的tag中 目前支持

```go
 where 查询条件
field 查询的字段
default 默认值
```
##实例
    
   ```go
import (
        "github.com/zngue/go_helper"
)
type MemberRequest struct {
    Username          int    `form:"username" field:"username" where:"eq" default:"-1"`
    ID                int    `form:"id" field:"id" where:"lt" default:"0"`
    Nickname          string `form:"nickname" field:"nickname" where:"like" default:""`
    CertificationName string `form:"certification_name" field:"certification_name" where:"like" default:""`
    Usernature        int    `form:"usernature" field:"usernature" where:"eq" default:"-1"`
    Usertype          int    `form:"usertype" field:"usertype" where:"eq" default:"-1"`
    Status            int    `form:"status" field:"status" where:"eq" default:"-1"`
    Stick             int    `form:"stick" field:"stick" where:"gt" default:"-1"`
    IDArr             string `form:"idArr" field:"id" where:"in" default:""`
    WhereOr	map[string]interface{} `form:"where_or"  where:"or"`
}
func main() {
    
	var rep MemberRequest
	//自己绑定自己数据 可以接受数据 也可以自己赋值传递
	///.........
	//调用实例
	//db gorm.DB 实例
	//rep  结构传入进去 如果传入值等于默认值 不参与where条件运算
	ext=&GormExt{DB:db,I:rep}
	ext.Init()
	
}

```

### where条件支持那些
<table>
<tbody>
    <tr>
        <th>表达式</th><th>文字说明</th><th>类型支持</th>
    </tr>
</tbody>
    <tr>
        <td>eq</td>
        <td>等于</td>
        <td>string|int</td>
    </tr>
<tr>
        <td>neq</td>
        <td>不等于</td>
<td>string|int</td>
    </tr>
<tr>
        <td>gt</td>
        <td>大于</td>
<td>string|int</td>
    </tr>
<tr>
        <td>egt</td>
        <td>大于等于</td>
        <td>string|int</td>

<tr>
        <td>lt</td>
        <td>小于</td>
        <td>string|int</td>
    </tr>
<tr>
        <td>elt</td>
        <td>小于等于</td>
        <td>string|int</td>
    </tr>
<tr>
        <td>null</td>
        <td>为空</td>
        <td>string|int</td>
    </tr>
<tr>
        <td>not</td>
        <td>不为空</td>
        <td>string|int</td>
    </tr>
<tr>
        <td>like</td>
        <td>模糊查询</td>
        <td>string</td>
    </tr>
<tr>
        <td>in</td>
        <td>条件in</td>
        <td>string,多个用逗号分隔</td>
    </tr>
<tr>
        <td>or</td>
        <td>或查询</td>
        <td>map[string]interface{}</td>
    </tr>

</table>
