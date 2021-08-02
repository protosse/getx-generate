import 'package:flutter/material.dart';
import 'home_controller.dart';
import '../../../routes/route_model.dart';
import 'package:get/get.dart';

class HomePage extends GetView<HomeController> {
	@override
	Widget build(Object context) {
		RouteModel rm = Get.arguments;
		return GetBuilder<HomeController>(
		  tag: rm.tag,
		  builder: (controller) {
			return Container();
		  },
		);
	}
}
	