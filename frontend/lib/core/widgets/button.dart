import 'package:flutter/material.dart';
import 'package:frontend/core/theme/app_colors.dart';

enum VRButtonType { primary, outlined }

class VRButton extends StatelessWidget {
  final IconData icon;
  final String label;
  final VoidCallback? onTap;
  final VRButtonType type;

  const VRButton({
    super.key,
    required this.icon,
    required this.label,
    required this.onTap,
    this.type = VRButtonType.primary,
  });

  bool get isOutlined => type == VRButtonType.outlined;

  @override
  Widget build(BuildContext context) {
    final Color primaryColor = AppColors.primary;
    final Color backgroundColor = isOutlined ? Colors.white : primaryColor;
    final Color textColor = isOutlined ? primaryColor : Colors.white;
    final Color borderColor = primaryColor;

    return InkWell(
      onTap: onTap ?? () {},
      borderRadius: BorderRadius.circular(20),
      child: Ink(
        width: double.infinity,
        padding: const EdgeInsets.symmetric(vertical: 16, horizontal: 24),
        decoration: BoxDecoration(
          color: backgroundColor,
          borderRadius: BorderRadius.circular(20),
          border: Border.all(color: borderColor, width: isOutlined ? 1.5 : 0),
          boxShadow: [
            if (!isOutlined)
              BoxShadow(
                // ignore: deprecated_member_use
                color: primaryColor.withOpacity(0.3),
                blurRadius: 8,
                offset: const Offset(0, 4),
              ),
          ],
        ),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Icon(icon, color: textColor, size: 20),
            const SizedBox(width: 12),
            Flexible(
              child: Text(
                label,
                style: TextStyle(
                  color: textColor,
                  fontSize: 16,
                  fontWeight: FontWeight.bold,
                ),
                overflow: TextOverflow.ellipsis,
              ),
            ),
          ],
        ),
      ),
    );
  }
}
