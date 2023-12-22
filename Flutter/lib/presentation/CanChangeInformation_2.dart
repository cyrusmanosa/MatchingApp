import 'package:dating_your_date/client/grpc_services.dart';
import 'package:dating_your_date/core/app_export.dart';
import 'package:dating_your_date/global_variable/model.dart';
import 'package:dating_your_date/pb/rpc_canChange.pb.dart';
import 'package:dating_your_date/widgets/app_bar/appbar_leading_image.dart';
import 'package:dating_your_date/widgets/app_bar/appbar_title.dart';
import 'package:dating_your_date/widgets/app_bar/custom_Input_Bar.dart';
import 'package:dating_your_date/widgets/app_bar/custom_app_bar.dart';
import 'package:dating_your_date/widgets/Custom_Outlined_Button.dart';
import 'package:dating_your_date/widgets/Custom_Input_Form_Bar.dart';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

// ignore_for_file: must_be_immutable,camel_case_types
class CanChangeInformation_2 extends StatelessWidget {
  CanChangeInformation_2({Key? key}) : super(key: key);
  TextEditingController canChangeJobController = TextEditingController();
  TextEditingController canChangeAnnualSalaryController = TextEditingController();
  TextEditingController canChangeSociabilityController = TextEditingController();
  TextEditingController canChangeReligiousController = TextEditingController();
  TextEditingController canChangeIntroduceController = TextEditingController();

  // Http
  void createCanChangeHttpRequest(BuildContext context) async {
    var url = "http://127.0.0.1:8080/CreateFixInfo";
    var requestBody = {
      "SessionID": globalSessionID,
      "NickName": nickname,
      "City": city,
      "Sexual": sexual,
      "Height": height,
      "Weight": weight,
      "Speaklanguage": speaklanguage,
      "Education": education,
      "Job": canChangeJobController,
      "AnnualSalary": int.parse(canChangeAnnualSalaryController.text),
      "Sociability": canChangeSociabilityController.text,
      "Religious": canChangeReligiousController.text,
      "Introduce": canChangeIntroduceController.text,
    };
    var response = await http.post(Uri.parse(url), body: jsonEncode(requestBody), headers: {"Content-Type": "application/json"});

    if (response.statusCode == 200) {
      onTapNextButton(context);
    } else {
      print("Job: ${canChangeJobController.text}");
      print("AnnualSalary: ${canChangeAnnualSalaryController.text}");
      print("Sociability: ${canChangeSociabilityController.text}");
      print("Religious: ${canChangeReligiousController.text}");
      print("Introduce: ${canChangeIntroduceController.text}");
    }
  }

  // Grpc
  void createCanChangeGrpcRequest(BuildContext context) async {
    final request = CreateCanChangeRequest(
      sessionID: globalSessionID,
      nickName: nickname,
      city: city,
      sexual: sexual,
      height: height,
      weight: weight,
      speaklanguage: speaklanguage,
      education: education,
      job: canChangeJobController.text,
      annualSalary: int.parse(canChangeAnnualSalaryController.text),
      sociability: canChangeSociabilityController.text,
      religious: canChangeReligiousController.text,
      introduce: canChangeIntroduceController.text,
    );

    final response = await GrpcService.client.createCanChange(request);
    // ignore: unnecessary_null_comparison
    if (response != null) {
      onTapNextButton(context);
    } else {
      showErrorDialog(context, "Error: validatable input data");
    }
  }

  void showErrorDialog(BuildContext context, String errorMessage) {
    showDialog(
      context: context,
      builder: (context) => AlertDialog(
        title: Text('Error'),
        content: Text(errorMessage),
        actions: [TextButton(onPressed: () => Navigator.pop(context), child: Text('OK'))],
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    mediaQueryData = MediaQuery.of(context);
    return SafeArea(
      child: Scaffold(
        appBar: _buildHeader(context),
        body: Padding(
          padding: EdgeInsets.symmetric(horizontal: mediaQueryData.size.width / 13, vertical: mediaQueryData.size.height / 20),
          child: Column(
            children: [
              // Job
              CustomInputBar(titleName: "仕事:", backendPart: _buildcanChangeJobInput(context)),
              SizedBox(height: mediaQueryData.size.height / 50),

              // Annual Salary
              CustomInputBar(titleName: "年収:", backendPart: _buildcanChangeAnnualSalaryInput(context)),
              SizedBox(height: mediaQueryData.size.height / 50),

              // Sociability
              CustomInputBar(titleName: "社交力:", backendPart: _buildcanChangeSociabilityInput(context)),
              SizedBox(height: mediaQueryData.size.height / 50),

              // Religious
              CustomInputBar(titleName: "宗教:", backendPart: _buildcanChangeReligiousInput(context)),
              SizedBox(height: mediaQueryData.size.height / 50),

              // Introduce
              CustomInputBar(titleName: "自己紹介:", backendPart: _buildcanChangeIntroduceInput(context)),

              SizedBox(height: 40),
              _buildNextPageButton(context),
            ],
          ),
        ),
      ),
    );
  }

  /// Header
  PreferredSizeWidget _buildHeader(BuildContext context) {
    return CustomAppBar(
      leading: AppbarLeadingImage(
        imagePath: ImageConstant.imgArrowLeft,
        margin: EdgeInsets.only(left: 25, top: 50, bottom: 10),
        onTap: () {
          onTapArrowLeft(context);
        },
      ),
      title: AppbarTitle(text: "基本個人情報 - C", margin: EdgeInsets.only(top: 60, bottom: 20)),
      styleType: Style.bgFill,
    );
  }

  // turn back
  onTapArrowLeft(BuildContext context) {
    Navigator.pop(context);
  }

  /// Job
  Widget _buildcanChangeJobInput(BuildContext context) {
    return CustomInputFormBar(
      controller: canChangeJobController,
      hintText: "ホスト",
      textInputAction: TextInputAction.done,
    );
  }

  /// Annual Salary
  Widget _buildcanChangeAnnualSalaryInput(BuildContext context) {
    return CustomInputFormBar(
      controller: canChangeAnnualSalaryController,
      hintText: "4000",
    );
  }

  /// Sociability
  Widget _buildcanChangeSociabilityInput(BuildContext context) {
    return CustomInputFormBar(
      controller: canChangeSociabilityController,
      hintText: "人たら神",
    );
  }

  /// Religious
  Widget _buildcanChangeReligiousInput(BuildContext context) {
    return CustomInputFormBar(
      controller: canChangeReligiousController,
      hintText: "多神教",
    );
  }

  /// Introduce
  Widget _buildcanChangeIntroduceInput(BuildContext context) {
    return CustomInputFormBar(
      controller: canChangeIntroduceController,
      hintText: "亜dさdさだだ",
      textInputAction: TextInputAction.done,
      maxLines: 8,
      contentPadding: EdgeInsets.symmetric(horizontal: 20.h, vertical: 20.v),
    );
  }

  /// Next Button
  Widget _buildNextPageButton(BuildContext context) {
    return CustomOutlinedButton(
      width: mediaQueryData.size.width / 4,
      height: mediaQueryData.size.height / 25,
      text: "次へ",
      buttonTextStyle: theme.textTheme.titleMedium,
      onPressed: () {
        createCanChangeGrpcRequest(context);
      },
    );
  }

  /// Navigates to the k26Screen when the action is triggered.
  onTapNextButton(BuildContext context) {
    Navigator.pushNamed(context, AppRoutes.searchTitle);
  }
}