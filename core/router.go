package core

type Route struct {
	Method string
	Path   string
}

//func RegisterRoutes(app *fiber.App, ctrl *MyController) {
//	// 通过反射扫描 MyController 中的方法
//	t := reflect.TypeOf(ctrl)
//	v := reflect.ValueOf(ctrl)
//
//	// 扫描每个方法并根据标签注册路由
//	for i := 0; i < t.NumMethod(); i++ {
//		method := t.Method(i)
//		fmt.Println(method)
//		// 获取每个方法的 "route" 标签，定义路由信息
//		//routeTag := method.Tag.Get("route")
//		//if routeTag != "" {
//		//	// 标签格式： "GET /home"
//		//	parts := strings.Split(routeTag, " ")
//		//	if len(parts) == 2 {
//		//		methodName := parts[0]
//		//		path := parts[1]
//		//
//		//		// 注册路由（这里只做了 GET 路由处理，其他方法类似）
//		//		if methodName == "GET" {
//		//			app.Get(path, func(c *fiber.Ctx) error {
//		//				return v.Method(i).Call([]reflect.Value{reflect.ValueOf(c)})[0].Interface().(error)
//		//			})
//		//		}
//		//	}
//		//}
//	}
//}
