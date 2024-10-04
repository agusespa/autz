// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Login() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html><head><meta charset=\"UTF-8\"><title>Login</title><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><script src=\"https://unpkg.com/htmx.org@1.9.10\"></script><style>\n        body {\n            font-family: 'Arial', sans-serif;\n            background-color: #000;\n            color: #fff;\n            display: flex;\n            justify-content: center;\n            align-items: center;\n            height: 100vh;\n            margin: 0;\n        }\n\n        form {\n            background-color: #000;\n            padding: 30px;\n            border: 1px solid #fff;\n            width: 100%;\n            max-width: 400px;\n            box-sizing: border-box;\n        }\n\n        h1 {\n            text-align: center;\n            margin-bottom: 20px;\n            color: #fff;\n            font-size: 24px;\n            font-weight: normal;\n        }\n\n        label {\n            display: block;\n            margin-bottom: 8px;\n            color: #fff;\n            font-size: 14px;\n            letter-spacing: 1px;\n        }\n\n        input {\n            width: 100%;\n            padding: 10px;\n            background-color: #000;\n            border: 1px solid #fff;\n            color: #fff;\n            margin-bottom: 20px;\n            font-size: 14px;\n            box-sizing: border-box;\n        }\n\n        input:focus {\n            border-color: #fff;\n            outline: none;\n        }\n\n        button {\n            width: 100%;\n            margin-top: 18px;\n            padding: 12px;\n            background-color: #000;\n            border: 1px solid #fff;\n            color: #fff;\n            font-size: 14px;\n            text-transform: uppercase;\n            letter-spacing: 2px;\n            cursor: pointer;\n            box-sizing: border-box;\n        }\n\n        button:hover {\n            background-color: #fff;\n            color: #000;\n        }\n\n        .indicator {\n            display: none;\n            font-size: 12px;\n            margin-left: 10px;\n        }\n\n        .htmx-request {\n            display: inline;\n        }\n\n        #message {\n            text-align: center;\n            padding: 20px 10px 0px 10px;\n            font-size: 14px;\n        }\n\n        .success {\n            color: #00ff00;\n        }\n\n        .error {\n            color: #ff0000;\n        }\n    </style><script>\n        const allowSwap = [403, 404];\n        document.addEventListener(\"DOMContentLoaded\", (event) => {\n            document.body.addEventListener('htmx:beforeSwap', function (evt) {\n                if (allowSwap.includes(evt.detail.xhr.status)) {\n                    evt.detail.shouldSwap = true;\n                    evt.detail.isError = false;\n                }\n            });\n        })\n    </script></head><body><form hx-post=\"/api/login\" hx-trigger=\"submit\" hx-target=\"#message\" hx-swap=\"innerHTML\" hx-headers=\"{&#34;Accept&#34;: &#34;application/json+cookie, text/html&#34;, &#34;X-Admin-Request&#34;: &#34;true&#34;}\" hx-indicator=\"#loading\" hx-disabled-elt=\"#loginButton\"><h1>Login</h1><div><label for=\"username\">Username</label> <input type=\"text\" id=\"username\" name=\"username\" placeholder=\"Enter username\" required></div><div><label for=\"password\">Password</label> <input type=\"password\" id=\"password\" name=\"password\" placeholder=\"Enter password\" required></div><button type=\"submit\" id=\"loginButton\">Log in <span id=\"loading\" class=\"indicator\">...</span></button><div id=\"message\"></div></form></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
