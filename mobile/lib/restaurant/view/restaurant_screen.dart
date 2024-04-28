import 'package:app/common/const/data.dart';
import 'package:app/restaurant/component/restaurant_card.dart';
import 'package:app/restaurant/model/restaurant_model.dart';
import 'package:app/restaurant/view/restaurant_detail_screen.dart';
import 'package:dio/dio.dart';
import 'package:flutter/material.dart';

class RestaurantScreen extends StatelessWidget {
  const RestaurantScreen({super.key});

  Future<List> paginateRestaurant() async {
    final dio = Dio();

    // Get Access token key
    final accessToken = await storage.read(key: ACCESS_TOKEN_KEY);
    // get data
    final resp = await dio.get('http://$ip/restaurant',
      options: Options(
        headers:{
          'authorization' : 'Bearer $accessToken',
        }
      )
    );

    // Return data
    return resp.data['data'];
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      child: Center(
        child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 16.0),
          child: FutureBuilder<List>(
            future: paginateRestaurant(),
            builder: (context, AsyncSnapshot<List> snapshot){
              if (!snapshot.hasData) {
                return Container(
                );
              }
              return ListView.separated(
                  itemBuilder: (_, index){
                    final item = snapshot.data![index];
                    final pItem = RestaurantModel.fromJson(json: item);
                    return GestureDetector(
                      onTap: (){
                        Navigator.of(context).push(
                          MaterialPageRoute(builder: (_)=>RestaurantDetailScreen()
                          )
                        );
                      },
                      child: RestaurantCard.fromModel(
                        model: pItem,
                      ),
                    );
                  },
                  // 아이템 사이사이 빌드
                  separatorBuilder: (_, index){
                    return const SizedBox(height: 16.0);
                  },
                  itemCount: snapshot.data!.length
              );
            },
          )
        ),
      ),
    );
  }
}
