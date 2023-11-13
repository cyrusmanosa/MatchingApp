import 'package:cyrus_man_s_application1/core/app_export.dart';
import 'package:cyrus_man_s_application1/widgets/app_bar/appbar_leading_image.dart';
import 'package:cyrus_man_s_application1/widgets/app_bar/appbar_title.dart';
import 'package:cyrus_man_s_application1/widgets/app_bar/custom_Input_Bar.dart';
import 'package:cyrus_man_s_application1/widgets/app_bar/custom_app_bar.dart';
import 'package:cyrus_man_s_application1/widgets/custom_outlined_button.dart';
import 'package:cyrus_man_s_application1/widgets/custom_text_form_field.dart';
import 'package:flutter/material.dart';

// ignore_for_file: must_be_immutable
class SignUp_2 extends StatelessWidget {
  SignUp_2({Key? key}) : super(key: key);

  TextEditingController basicNickNameInputController = TextEditingController();
  TextEditingController basicCityInputController = TextEditingController();
  TextEditingController basicSexualInputController = TextEditingController();
  TextEditingController basicHeightInputController = TextEditingController();
  TextEditingController basicWeightInputController = TextEditingController();
  TextEditingController basicEducationInputController = TextEditingController();
  TextEditingController basicJobInputController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    mediaQueryData = MediaQuery.of(context);
    return SafeArea(
      child: Scaffold(
        body: SizedBox(
          width: double.maxFinite,
          child: SingleChildScrollView(
            child: Column(
              children: [
                // Header
                _buildHeader(context),
                SizedBox(height: 50),

                // Nick Name
                CustomInputBar(titleName: "ニックネーム:", backendPart: _buildBasicNickNameInput(context)),
                SizedBox(height: 15.v),

                // City
                CustomInputBar(titleName: "居住地:", backendPart: _buildBasicCityInput(context)),
                SizedBox(height: 15.v),

                // Sexual
                CustomInputBar(titleName: "性的指向:", backendPart: _buildBasicSexualInput(context)),
                SizedBox(height: 15.v),

                // Height
                CustomInputBar(titleName: "身長:", backendPart: _buildBasicHeightInput(context)),
                SizedBox(height: 15.v),

                // Weight
                CustomInputBar(titleName: "体重:", backendPart: _buildBasicWeightInput(context)),
                SizedBox(height: 15.v),

                // Education
                CustomInputBar(titleName: "学歴:", backendPart: _buildBasicEducationInput(context)),
                SizedBox(height: 15.v),

                // Job
                CustomInputBar(titleName: "職種:", backendPart: _buildBasicJobInput(context)),
                SizedBox(height: 40.v),

                // Button
                _buildNextButton(context),
                SizedBox(height: 30.v),
              ],
            ),
          ),
        ),
      ),
    );
  }

// Backend ----------------------------------------------------------------

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
      title: AppbarTitle(text: "基本個人情報 - B", margin: EdgeInsets.only(top: 60, bottom: 20)),
      styleType: Style.bgFill,
    );
  }

  // turn back
  onTapArrowLeft(BuildContext context) {
    Navigator.pop(context);
  }

  /// Nickname
  Widget _buildBasicNickNameInput(BuildContext context) {
    return CustomTextFormField(
      controller: basicNickNameInputController,
      hintText: "仆街",
    );
  }

  /// City
  Widget _buildBasicCityInput(BuildContext context) {
    return CustomTextFormField(
      controller: basicCityInputController,
      hintText: "大阪",
    );
  }

  /// Sexual
  Widget _buildBasicSexualInput(BuildContext context) {
    return CustomTextFormField(
      controller: basicSexualInputController,
      hintText: "異性愛",
    );
  }

  /// Height
  Widget _buildBasicHeightInput(BuildContext context) {
    return CustomTextFormField(
      controller: basicHeightInputController,
      hintText: "170cm",
    );
  }

  /// Weight
  Widget _buildBasicWeightInput(BuildContext context) {
    return CustomTextFormField(
      controller: basicWeightInputController,
      hintText: "60kg",
    );
  }

  /// Education
  Widget _buildBasicEducationInput(BuildContext context) {
    return CustomTextFormField(
      controller: basicEducationInputController,
      hintText: "高校生",
    );
  }

  /// Job
  Widget _buildBasicJobInput(BuildContext context) {
    return CustomTextFormField(
      controller: basicJobInputController,
      hintText: "ホスト",
      textInputAction: TextInputAction.done,
    );
  }

  /// Section Widget
  Widget _buildNextButton(BuildContext context) {
    return CustomOutlinedButton(
      width: 150.h,
      text: "次へ",
      margin: EdgeInsets.only(left: 140.h, right: 140.h, bottom: 30.v),
      buttonStyle: CustomButtonStyles.outlinePinkGrayBG,
      onPressed: () {
        onTapNextButton(context);
      },
    );
  }

  /// Navigates to the signupPhoneoremailPartthreeScreen when the action is triggered.
  onTapNextButton(BuildContext context) {
    Navigator.pushNamed(context, AppRoutes.signUp_3);
  }
}
