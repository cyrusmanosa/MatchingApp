import 'package:dating_your_date/core/app_export.dart';
import 'package:flutter/material.dart';

class TargetReady extends StatelessWidget {
  TargetReady({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    bool isPageNavigated = false;
    Future.delayed(
      Duration(seconds: 3),
      () {
        if (!isPageNavigated) {
          isPageNavigated = true;
          Navigator.pushNamed(context, AppRoutes.containerScreen);
        }
      },
    );

    MediaQueryData mediaQueryData = MediaQuery.of(context);
    double mediaH = mediaQueryData.size.height;
    double mediaW = mediaQueryData.size.width;

    return Scaffold(
      body: Container(
        padding: EdgeInsets.symmetric(horizontal: mediaW / 13, vertical: mediaH / 20),
        child: Align(
          child: Column(
            children: [
              // Logo and Slogan
              SizedBox(height: mediaH / 15),
              CustomImageView(imagePath: ImageConstant.imgLogo, width: mediaW / 3.5),
              CustomImageView(imagePath: ImageConstant.imgSlogan, width: mediaW / 3),
              SizedBox(height: mediaH / 15),

              // Title
              Text("ターゲットを準備します", style: TextStyle(fontSize: 25, color: appTheme.black)),
            ],
          ),
        ),
      ),
    );
  }
}
