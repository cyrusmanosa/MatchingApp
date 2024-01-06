import 'package:dating_your_date/client/grpc_services.dart';
import 'package:dating_your_date/core/app_export.dart';
import 'package:dating_your_date/models/model.dart';
import 'package:dating_your_date/pb/canChange.pb.dart';
import 'package:dating_your_date/pb/rpc_canChange.pb.dart';
import 'package:dating_your_date/widgets/app_bar/appbar_title.dart';
import 'package:dating_your_date/widgets/app_bar/custom_Input_bar.dart';
import 'package:dating_your_date/widgets/Custom_Outlined_Button.dart';
import 'package:dating_your_date/widgets/Custom_Input_Form_Bar.dart';
import 'package:dating_your_date/widgets/Custom_WarningMsgBox.dart';
import 'package:flutter/material.dart';
import 'package:grpc/grpc.dart';

class ProfileEdit extends StatefulWidget {
  ProfileEdit({Key? key, this.canData}) : super(key: key);
  final CanChange? canData;

  @override
  _ProfileEditState createState() => _ProfileEditState();
}

class _ProfileEditState extends State<ProfileEdit> {
  TextEditingController updateNickNameController = TextEditingController();
  TextEditingController updateCityController = TextEditingController();
  TextEditingController updateSexualController = TextEditingController();
  TextEditingController updateHeightController = TextEditingController();
  TextEditingController updateWeightController = TextEditingController();
  TextEditingController updateEducationController = TextEditingController();
  TextEditingController updateJobController = TextEditingController();
  TextEditingController updateSpeakLanguageController = TextEditingController();
  TextEditingController updateAnnualSalaryController = TextEditingController();
  TextEditingController updateSociabilityController = TextEditingController();
  TextEditingController updateReligiousController = TextEditingController();
  TextEditingController updateIntroduceController = TextEditingController();

  @override
  void initState() {
    super.initState();
    updateNickNameController = TextEditingController(text: widget.canData?.nickName);
    updateCityController = TextEditingController(text: widget.canData?.city);
    updateSexualController = TextEditingController(text: widget.canData?.sexual);
    updateHeightController = TextEditingController(text: widget.canData?.height.toString());
    updateWeightController = TextEditingController(text: widget.canData?.weight.toString());
    updateEducationController = TextEditingController(text: widget.canData?.education);
    updateSpeakLanguageController = TextEditingController(text: widget.canData?.speaklanguage);
    updateJobController = TextEditingController(text: widget.canData?.job);
    updateAnnualSalaryController = TextEditingController(text: widget.canData?.annualSalary.toString());
    updateSociabilityController = TextEditingController(text: widget.canData?.sociability);
    updateReligiousController = TextEditingController(text: widget.canData?.religious);
    updateIntroduceController = TextEditingController(text: widget.canData?.introduce);
  }

  bool isPureNumber(String value) {
    final pattern = RegExp(r'^\d+$');
    return pattern.hasMatch(value);
  }

  // Grpc
  void updateDataGrpcRequest(BuildContext context) async {
    if (updateNickNameController.text.isEmpty) {
      showErrorDialog(context, "ニックネームはまだ入力されていません");
    } else if (updateCityController.text.isEmpty) {
      showErrorDialog(context, "居住地はまだ入力されていません");
    } else if (updateSexualController.text.isEmpty) {
      showErrorDialog(context, "性的指向はまだ入力されていません");
    } else if (updateHeightController.text.isEmpty) {
      showErrorDialog(context, "身長はまだ入力されていません");
    } else if (!isPureNumber(updateHeightController.text)) {
      showErrorDialog(context, "入力した身長は数字じゃありません");
    } else if (updateWeightController.text.isEmpty) {
      showErrorDialog(context, "体重はまだ入力されていません");
    } else if (!isPureNumber(updateWeightController.text)) {
      showErrorDialog(context, "入力した体重は数字じゃありません");
    } else if (updateSpeakLanguageController.text.isEmpty) {
      showErrorDialog(context, "学歴はまだ入力されていません");
    } else if (updateJobController.text.isEmpty) {
      showErrorDialog(context, "仕事はまだ入力されていません");
    } else if (!isPureNumber(updateAnnualSalaryController.text)) {
      showErrorDialog(context, "入力した年収は数字じゃありません");
    } else if (updateSociabilityController.text.isEmpty) {
      showErrorDialog(context, "社交力はまだ入力されていません");
    } else if (updateReligiousController.text.isEmpty) {
      showErrorDialog(context, "宗教はまだ入力されていません");
    } else if (updateIntroduceController.text.isEmpty) {
      showErrorDialog(context, "自己紹介はまだ入力されていません");
    } else {
      String? apiKeyS = await globalSession.read(key: 'SessionId');
      final request = UpdateCanChangeRequest(
        sessionID: apiKeyS,
        nickName: updateNickNameController.text,
        city: updateCityController.text,
        sexual: updateSexualController.text,
        height: int.parse(updateHeightController.text),
        weight: int.parse(updateWeightController.text),
        speaklanguage: updateSpeakLanguageController.text,
        education: updateEducationController.text,
        job: updateJobController.text,
        annualSalary: int.parse(updateAnnualSalaryController.text),
        sociability: updateSociabilityController.text,
        religious: updateReligiousController.text,
        introduce: updateIntroduceController.text,
      );

      try {
        await GrpcInfoService.client.updateCanChange(request);
        onTapNextPage(context);
        showDoneDialog(context, "個人情報もアップしました");
      } on GrpcError catch (e) {
        showErrorDialog(context, "Error: ${e.codeName}");
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    MediaQueryData mediaQueryData = MediaQuery.of(context);
    return Scaffold(
      appBar: _buildHeader(context),
      resizeToAvoidBottomInset: false,
      body: Container(
        width: double.maxFinite,
        child: SingleChildScrollView(
          padding: EdgeInsets.symmetric(horizontal: mediaQueryData.size.width / 13, vertical: mediaQueryData.size.height / 20),
          child: Column(
            children: [
              // photos
              Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  CustomImageView(
                    imagePath: ImageConstant.imgVectorgray500,
                    height: 150,
                    width: 150,
                    margin: EdgeInsets.symmetric(vertical: 10),
                  ),
                  CustomImageView(
                    imagePath: ImageConstant.imgPhotoSet,
                    height: 170,
                    width: 50,
                    margin: EdgeInsets.only(left: 22),
                  ),
                ],
              ),

              // Introduce
              CustomInputBar(titleName: "自己紹介:", backendPart: _buildUpdateIntroduceInput(context)),
              SizedBox(height: mediaQueryData.size.height / 50),

              // Nickname
              CustomInputBar(titleName: "ニックネーム:", backendPart: _buildUpdateNickNameInput(context)),
              SizedBox(height: mediaQueryData.size.height / 50),

              // Height
              CustomInputBar(titleName: "身長:", backendPart: _buildUpdateHeightInput(context)),
              SizedBox(height: mediaQueryData.size.height / 50),

              // Weight
              CustomInputBar(titleName: "体重:", backendPart: _buildUpdateWeightInput(context)),
              SizedBox(height: mediaQueryData.size.height / 50),

              // City
              CustomInputBar(titleName: "居住地:", backendPart: _buildUpdateCityInput(context)),
              SizedBox(height: mediaQueryData.size.height / 50),

              // Education
              CustomInputBar(titleName: "学歴:", backendPart: _buildUpdateEducationInput(context)),
              SizedBox(height: mediaQueryData.size.height / 50),

              // Speak Language
              CustomInputBar(titleName: "言語:", backendPart: _buildUpdateSpeakLanguageInput(context)),
              SizedBox(height: mediaQueryData.size.height / 50),

              // Job
              CustomInputBar(titleName: "職種:", backendPart: _buildUpdateJobInput(context)),
              SizedBox(height: mediaQueryData.size.height / 50),

              // Annual Salary
              CustomInputBar(titleName: "年収:", backendPart: _buildUpdateAnnualSalaryInput(context)),
              SizedBox(height: mediaQueryData.size.height / 50),

              // Sexual
              CustomInputBar(titleName: "性的指向:", backendPart: _buildUpdateSexualInput(context)),
              SizedBox(height: mediaQueryData.size.height / 50),

              // Sociability
              CustomInputBar(titleName: "社交力:", backendPart: _buildUpdateSociabilityInput(context)),
              SizedBox(height: mediaQueryData.size.height / 50),

              // Relighious
              CustomInputBar(titleName: "宗教:", backendPart: _buildUpdateReligiousInput(context)),
              SizedBox(height: mediaQueryData.size.height / 25),

              // button
              _buildSubmitButton(context),
            ],
          ),
        ),
      ),
    );
  }

  /// Header
  PreferredSizeWidget _buildHeader(BuildContext context) {
    return AppBar(automaticallyImplyLeading: true, title: AppbarTitle(text: "プロフィール編集"));
  }

  /// Nickname
  Widget _buildUpdateNickNameInput(BuildContext context) {
    return CustomInputFormBar(controller: updateNickNameController, hintText: "仆街");
  }

  /// City
  Widget _buildUpdateCityInput(BuildContext context) {
    return CustomInputFormBar(controller: updateCityController, hintText: "大阪");
  }

  /// Sexual
  Widget _buildUpdateSexualInput(BuildContext context) {
    return CustomInputFormBar(controller: updateSexualController, hintText: "異性愛");
  }

  /// Height
  Widget _buildUpdateHeightInput(BuildContext context) {
    return CustomInputFormBar(controller: updateHeightController, hintText: "170cm");
  }

  /// Weight
  Widget _buildUpdateWeightInput(BuildContext context) {
    return CustomInputFormBar(controller: updateWeightController, hintText: "60kg");
  }

  /// Education
  Widget _buildUpdateEducationInput(BuildContext context) {
    return CustomInputFormBar(controller: updateEducationController, hintText: "高校生");
  }

  /// Speak Language
  Widget _buildUpdateSpeakLanguageInput(BuildContext context) {
    return CustomInputFormBar(controller: updateSpeakLanguageController, hintText: "言語");
  }

  /// Job
  Widget _buildUpdateJobInput(BuildContext context) {
    return CustomInputFormBar(controller: updateJobController, hintText: "ホスト");
  }

  /// Annual Salary
  Widget _buildUpdateAnnualSalaryInput(BuildContext context) {
    return CustomInputFormBar(controller: updateAnnualSalaryController, hintText: "4000");
  }

  /// Sociability
  Widget _buildUpdateSociabilityInput(BuildContext context) {
    return CustomInputFormBar(controller: updateSociabilityController, hintText: "人たら神");
  }

  /// Religious
  Widget _buildUpdateReligiousInput(BuildContext context) {
    return CustomInputFormBar(controller: updateReligiousController, hintText: "多神教");
  }

  /// Introduce
  Widget _buildUpdateIntroduceInput(BuildContext context) {
    MediaQueryData mediaQueryData = MediaQuery.of(context);
    return CustomInputFormBar(
      prefix: Padding(padding: EdgeInsets.symmetric(horizontal: mediaQueryData.size.width / 75)),
      height: mediaQueryData.size.height / 5,
      controller: updateIntroduceController,
      hintText: "亜dさdさだだ",
      maxLines: 8,
      focusNode: FocusNode(),
      onTap: () {
        FocusNode().requestFocus();
      },
      contentPadding: EdgeInsets.symmetric(horizontal: mediaQueryData.size.height / 200, vertical: mediaQueryData.size.width / 50),
    );
  }

  /// button
  Widget _buildSubmitButton(BuildContext context) {
    return CustomOutlinedButton(
      text: "送信",
      onPressed: () {
        updateDataGrpcRequest(context);
      },
    );
  }

  onTapNextPage(BuildContext context) {
    Navigator.pop(context);
  }

  void showDoneDialog(BuildContext context, String s) {}
}
