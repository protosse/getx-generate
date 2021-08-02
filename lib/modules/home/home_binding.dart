import 'package:get/get.dart';
import '../../../routes/route_model.dart';
import 'home_controller.dart';

class HomeBinding implements Bindings {
	@override
	void dependencies() {
		RouteModel rm = Get.arguments;
		var controller = HomeController(param: rm.param);
		Get.lazyPut<HomeController>(
		  () => controller,
		  tag: rm.tag,
		);
	}
}
	