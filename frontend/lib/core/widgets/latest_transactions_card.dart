import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:frontend/core/services/snackbar_service.dart';

class LatestTransactionsWidget extends StatelessWidget {
  final List<Map<String, dynamic>> transactions;

  const LatestTransactionsWidget({super.key, required this.transactions});

  @override
  Widget build(BuildContext context) {
    if (transactions.isEmpty) return const SizedBox.shrink();

    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        const SizedBox(height: 32),
        const Text(
          'Últimas 5 Transações',
          style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold),
        ),
        const SizedBox(height: 16),
        ...transactions.map(
          (tx) => Container(
            margin: const EdgeInsets.only(bottom: 12),
            padding: const EdgeInsets.symmetric(vertical: 12, horizontal: 16),
            decoration: BoxDecoration(
              color: Colors.white,
              border: Border.all(color: Colors.grey.shade300),
              borderRadius: BorderRadius.circular(12),
              boxShadow: [
                BoxShadow(
                  color: Colors.grey.shade200,
                  blurRadius: 4,
                  offset: const Offset(0, 2),
                ),
              ],
            ),
            child: Row(
              children: [
                Expanded(
                  child: Text(
                    '${tx['description']} • ${(DateTime.parse(tx['date'])).toIso8601String().split("T").first} • \$${(tx['amount_usd'] as num).toStringAsFixed(2)} • ${tx['id']}',
                    style: const TextStyle(fontSize: 13),
                    overflow: TextOverflow.ellipsis,
                  ),
                ),
                IconButton(
                  icon: const Icon(Icons.copy, size: 18),
                  tooltip: 'Copiar UUID',
                  onPressed: () async {
                    await Clipboard.setData(ClipboardData(text: tx['id']));
                    if (!context.mounted) return;
                    SnackBarService.showSuccess(context, 'UUID copiado!');
                  },
                ),
              ],
            ),
          ),
        ),
      ],
    );
  }
}
