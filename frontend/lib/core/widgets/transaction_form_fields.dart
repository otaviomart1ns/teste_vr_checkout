import 'package:flutter/material.dart';
import 'package:frontend/core/utils/input_formatters.dart';
import 'package:frontend/core/utils/validators.dart';

class TransactionFormFields extends StatelessWidget {
  final TextEditingController descriptionController;
  final TextEditingController valueController;
  final DateTime selectedDate;
  final void Function(DateTime newDate) onDateChanged;

  const TransactionFormFields({
    super.key,
    required this.descriptionController,
    required this.valueController,
    required this.selectedDate,
    required this.onDateChanged,
  });

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        // Descrição
        TextFormField(
          key: const Key('description-field'),
          controller: descriptionController,
          maxLength: 50,
          decoration: const InputDecoration(
            labelText: 'Descrição',
            border: OutlineInputBorder(),
            counterText: '',
          ),
          validator: Validators.validateDescription,
        ),
        const SizedBox(height: 16),

        GestureDetector(
          key: const Key('date-field'),
          onTap: () async {
            final today = DateTime.now();
            final picked = await showDatePicker(
              context: context,
              initialDate: selectedDate.isAfter(today) ? today : selectedDate,
              firstDate: DateTime(today.year - 5, today.month, today.day),
              lastDate: today,
            );
            if (picked != null) {
              onDateChanged(picked);
            }
          },
          child: InputDecorator(
            decoration: const InputDecoration(
              labelText: 'Data da Transação',
              border: OutlineInputBorder(),
            ),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Text('${selectedDate.toLocal()}'.split(' ')[0]),
                const Icon(Icons.calendar_today, size: 20),
              ],
            ),
          ),
        ),
        const SizedBox(height: 4),

        Builder(
          builder: (context) {
            final error = Validators.validateTransactionDate(selectedDate);
            if (error != null) {
              return Padding(
                padding: const EdgeInsets.only(top: 4),
                child: Align(
                  alignment: Alignment.centerLeft,
                  child: Text(
                    error,
                    style: TextStyle(
                      color: Theme.of(context).colorScheme.error,
                      fontSize: 12,
                    ),
                  ),
                ),
              );
            }
            return const SizedBox.shrink();
          },
        ),
        const SizedBox(height: 16),

        TextFormField(
          key: const Key('value-field'),
          controller: valueController,
          keyboardType: const TextInputType.numberWithOptions(decimal: true),
          inputFormatters: [InputFormatters.currencyFormatter()],
          decoration: const InputDecoration(
            labelText: 'Valor da Compra (USD)',
            border: OutlineInputBorder(),
          ),
          validator: Validators.validateValue,
        ),
      ],
    );
  }
}
