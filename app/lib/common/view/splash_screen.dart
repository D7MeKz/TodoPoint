import 'package:dio/dio.dart';
import 'package:flutter/material.dart';

import '../const/data.dart';

class SplashScreen extends StatefulWidget {
  const SplashScreen({super.key});

  @override
  State<SplashScreen> createState() => _SplashScreenState();
}

class _SplashScreenState extends State<SplashScreen> {
  @override
  Widget build(BuildContext context) {
    return const Placeholder();
  }
  @override
  void initState() {
    super.initState();
    //deleteToken();
    checkToken();
  }


  void deleteToken() async{
    await storage.deleteAll();
  }

  void checkToken() async{
    final refreshToken = await storage.read(key:REFRESH_TOKEN_KEY);
    final accessToken = await storage.read(key:ACCESS_TOKNE_KEY);
    final dio = Dio();

    // // TODO Validate token
    // try{
    //   final resp = await dio.post('http://localhost:3000/auth/token',
    //       data: {'refresh_token': refreshToken}
    //   );
    //
    //   // Go to Root Tab
    //   Navigator.of(context).pushAndRemoveUntil(
    //       MaterialPageRoute(
    //         builder: (_) => const RootTab(),
    //       ), (route) => false);
    // }catch(e){
    //   // Go to login screen
    //   Navigator.of(context).pushAndRemoveUntil(
    //       MaterialPageRoute(
    //         builder: (_) => const LoginScreen(),
    //       ), (route) => false);
    // }

  }
}