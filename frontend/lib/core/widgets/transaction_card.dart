import 'package:flutter/material.dart';
import 'package:frontend/core/theme/app_colors.dart';

class TransactionCard extends StatelessWidget {
  final String description;
  final DateTime date;
  final double originalValue;
  final double? convertedValue;
  final double? exchangeRate;
  final List<Widget>? actions;

  const TransactionCard({
    super.key,
    required this.description,
    required this.date,
    required this.originalValue,
    this.convertedValue,
    this.exchangeRate,
    this.actions,
  });

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);

    return Card(
      elevation: 0,
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(16),
        // ignore: deprecated_member_use
        side: BorderSide(color: AppColors.primary.withOpacity(0.2)),
      ),
      // ignore: deprecated_member_use
      color: AppColors.primary.withOpacity(0.05),
      margin: const EdgeInsets.only(bottom: 16),
      child: Padding(
        padding: const EdgeInsets.all(20),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            // Título
            Text(
              description,
              style: theme.textTheme.titleLarge?.copyWith(
                fontWeight: FontWeight.bold,
              ),
            ),
            const SizedBox(height: 16),

            // Data
            Row(
              children: [
                const Icon(Icons.calendar_today, size: 20),
                const SizedBox(width: 8),
                Text('Data: ${date.toLocal().toString().split(' ')[0]}'),
              ],
            ),
            const SizedBox(height: 12),

            // Valor Original
            Row(
              children: [
                const Icon(Icons.attach_money, size: 20),
                const SizedBox(width: 8),
                Text('Valor: \$${originalValue.toStringAsFixed(2)}'),
              ],
            ),

            if (convertedValue != null && exchangeRate != null) ...[
              const SizedBox(height: 12),

              // Valor Convertido
              Row(
                children: [
                  const Icon(Icons.currency_exchange, size: 20),
                  const SizedBox(width: 8),
                  Text(
                    'Valor Convertido: \$${convertedValue!.toStringAsFixed(2)}',
                  ),
                ],
              ),
              const SizedBox(height: 12),

              // Taxa de Câmbio
              Row(
                children: [
                  const Icon(Icons.compare_arrows, size: 20),
                  const SizedBox(width: 8),
                  Text('Taxa de Câmbio: ${exchangeRate!.toStringAsFixed(2)}'),
                ],
              ),
            ],

            // Ações
            if (actions != null && actions!.isNotEmpty) ...[
              const SizedBox(height: 16),
              Row(mainAxisAlignment: MainAxisAlignment.end, children: actions!),
            ],
          ],
        ),
      ),
    );
  }
}
