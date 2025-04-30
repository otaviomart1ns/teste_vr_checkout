import 'package:flutter/material.dart';

class SnackBarService {
  static void showSuccess(BuildContext context, String message) {
    _show(context, message, backgroundColor: Colors.green);
  }

  static void showError(BuildContext context, String message) {
    _show(context, message, backgroundColor: Colors.red);
  }

  static void showInfo(BuildContext context, String message) {
    _show(context, message, backgroundColor: Colors.yellow);
  }

  static void _show(
    BuildContext context,
    String message, {
    required Color backgroundColor,
  }) {
    final messenger = ScaffoldMessenger.of(
      Navigator.of(context, rootNavigator: true).context,
    );
    messenger.clearSnackBars();
    messenger.showSnackBar(
      SnackBar(
        content: Row(
          children: [
            const Icon(Icons.info_outline, color: Colors.white, size: 20),
            const SizedBox(width: 8),
            Expanded(
              child: Text(message, style: const TextStyle(fontSize: 14)),
            ),
          ],
        ),
        backgroundColor: backgroundColor,
        behavior: SnackBarBehavior.floating,
        margin: const EdgeInsets.symmetric(horizontal: 24, vertical: 16),
        shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(12)),
        duration: const Duration(milliseconds: 2500),
        showCloseIcon: true,
        closeIconColor: Colors.white,
        dismissDirection: DismissDirection.horizontal,
        elevation: 8,
      ),
    );
  }
}
