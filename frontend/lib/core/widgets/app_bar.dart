import 'package:flutter/material.dart';
import 'package:frontend/core/theme/app_colors.dart';

class VRAppBar extends StatelessWidget implements PreferredSizeWidget {
  const VRAppBar({super.key});

  @override
  Widget build(BuildContext context) {
    return AppBar(
      // ignore: deprecated_member_use
      backgroundColor: AppColors.primary.withOpacity(0.05),
      elevation: 0,
      centerTitle: true,
      automaticallyImplyLeading: false,
      title: Padding(
        padding: const EdgeInsets.symmetric(vertical: 8.0),
        child: Image.asset('assets/images/logo_vr.png', height: 40),
      ),
    );
  }

  @override
  Size get preferredSize => const Size.fromHeight(70);
}
