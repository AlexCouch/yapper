// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Sidebar() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"m-5 lg:m-4 mt-3 border-r border-r-gray-900 dark:border-r-gray-600\"><div class=\"flex flex-col items-end\"><!-- Basic user info --><div class=\"w-64 shrink-0 basis-auto flex flex-col items-start\"><!-- Home --><a href=\"/\" class=\"cursor-pointer flex flex-col items-start flex-1 w-full text-amber-700 hover:bg-amber-700/[.5] transition duration-200 dark:text-white hover:text-white dark:hover:bg-sky-700/[.50] py-4 px-2\"><div class=\"cursor-pointer flex flex-row justify-center items-center space-x-5 \"><div class=\"fa-solid fa-home\" style=\"font-size: 2.5em;\"></div><div class=\"rounded-lg mx-w-md w-full\"><h2 class=\"text-xl font-bold\">Home</h2></div></div></a><!-- TODO: Add route based on provided user id --><a class=\"cursor-pointer flex flex-col items-start flex-1 w-full text-amber-700 hover:bg-amber-700 transition duration-200 dark:text-white hover:text-white dark:hover:bg-sky-700/[.50] py-4 px-2\"><div class=\"cursor-pointer flex flex-row justify-center items-center space-x-5 \"><img src=\"pfp.png\" alt=\"Profile Picture\" class=\"bottom-0 w-10 h-10\"><div class=\"mx-w-md w-full flex flex-col justify-center items-center\"><!-- User Info with Verified Button --><h2 class=\"text-sm font-bold \">Alex Couch</h2><!-- Bio --><p class=\"dark:text-gray-300 text-sm\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs("@alexcouch")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `cmd/views/content.templ`, Line: 30, Col: 22}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></div></div></a><!-- Notifications --><!-- Add bubble for number of notifications unread --><a class=\"cursor-pointer flex flex-col items-start flex-1 w-full text-amber-700 hover:bg-amber-700 transition duration-200 dark:text-white hover:text-white dark:hover:bg-sky-700/[.50] py-4 px-2\"><div class=\"cursor-pointer flex flex-row justify-center items-center space-x-5 \"><div class=\"relative\"><span class=\"fa fa-envelope\" style=\"font-size: 2.5em;\"></span> <span class=\"absolute bg-sky-700 text-white px-2 py-1 text-xs font-bold rounded-full -bottom-2 -right-2\">99+</span></div><div class=\"rounded-lg mx-w-md w-full\"><h2 class=\"text-xl font-bold\">Notifications</h2></div></div></a><!-- Profile --></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}