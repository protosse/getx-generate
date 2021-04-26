import 'package:get/get.dart';
import 'search_list_controller.dart';

class SearchListBinding implements Bindings {
  @override
  void dependencies() {
    Get.lazyPut<SearchListController>(() => SearchListController());
  }
}
		