import 'package:cyrus_man_s_application1/core/app_export.dart';
import 'package:cyrus_man_s_application1/widgets/custom_elevated_button.dart';
import 'package:flutter/material.dart';

// ignore: must_be_immutable
class ResetTargetButtonLover extends StatelessWidget {
  ResetTargetButtonLover({Key? key}) : super(key: key);
  @override
  Widget build(BuildContext context) {
    return CustomElevatedButton(
      width: 350.h,
      text: "恋人",
      buttonStyle: CustomButtonStyles.fillDarkRed,
      onPressed: () {
        onTaptf(context);
      },
    );
  }
}

/// Navigates to the k3Screen when the action is triggered.
onTaptf(BuildContext context) {
  Navigator.pushNamed(context, AppRoutes.loverConditionsRepair);
}
