import 'package:json_annotation/json_annotation.dart';
part 'home.g.dart';

@JsonSerializable()
class Home {
	Home({});

	factory Home.fromJson(Map<String, dynamic> json) => _$HomeFromJson(json);
	Map<String, dynamic> toJson() => _$HomeToJson(this);
}		
	