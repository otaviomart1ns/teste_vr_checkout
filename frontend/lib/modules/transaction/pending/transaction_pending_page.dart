import 'package:flutter/material.dart';
import 'package:flutter_mobx/flutter_mobx.dart';
import 'package:flutter_modular/flutter_modular.dart';
import 'package:flutter_multi_formatter/formatters/formatter_utils.dart';
import 'package:frontend/core/services/snackbar_service.dart';
import 'package:frontend/core/widgets/app_bar.dart';
import 'package:frontend/core/widgets/button.dart';
import 'package:frontend/core/widgets/transaction_card.dart';
import 'package:frontend/core/widgets/transaction_form_fields.dart';
import 'package:frontend/modules/transaction/pending/store/transaction_pending_store.dart';

class TransactionPendingPage extends StatefulWidget {
  const TransactionPendingPage({super.key});

  @override
  State<TransactionPendingPage> createState() => _TransactionPendingPageState();
}

class _TransactionPendingPageState extends State<TransactionPendingPage> {
  late final TransactionPendingStore store;

  @override
  void initState() {
    super.initState();
    store = Modular.get<TransactionPendingStore>();
    store.loadPendingTransactions();
  }

  void _editTransaction(BuildContext context, String id) {
    final transaction = store.pendingTransactions.firstWhere((t) => t.id == id);

    final formKey = GlobalKey<FormState>();
    final descriptionController = TextEditingController(
      text: transaction.description,
    );
    final valueController = TextEditingController(
      text: transaction.amountUsd.toStringAsFixed(2),
    );
    DateTime selectedDate = transaction.date;

    showDialog(
      context: context,
      builder: (context) => StatefulBuilder(
        builder: (context, setStateDialog) => AlertDialog(
          title: const Text('Editar Transação'),
          content: Form(
            key: formKey,
            autovalidateMode: AutovalidateMode.onUserInteraction,
            child: SingleChildScrollView(
              child: TransactionFormFields(
                descriptionController: descriptionController,
                valueController: valueController,
                selectedDate: selectedDate,
                onDateChanged: (newDate) {
                  setStateDialog(() {
                    selectedDate = newDate;
                  });
                },
              ),
            ),
          ),
          actions: [
            Row(
              children: [
                Expanded(
                  child: VRButton(
                    icon: Icons.cancel,
                    label: 'Cancelar',
                    onTap: () => Modular.to.pop(),
                    type: VRButtonType.outlined,
                  ),
                ),
                const SizedBox(width: 8),
                Expanded(
                  child: VRButton(
                    icon: Icons.save,
                    label: 'Salvar',
                    onTap: () async {
                      if (!(formKey.currentState?.validate() ?? false)) {
                        return;
                      }

                      final cleanDescription = descriptionController.text
                          .trim();
                      final cleanValue = toNumericString(
                        valueController.text,
                        allowPeriod: true,
                      );
                      final double number = double.parse(cleanValue);

                      await store.editPendingTransaction(
                        id: id,
                        newDescription: cleanDescription,
                        newDate: selectedDate,
                        newAmountUsd: number,
                      );

                      if (store.errorMessage != null) {
                        // ignore: use_build_context_synchronously
                        SnackBarService.showError(context, store.errorMessage!);
                      } else {
                        SnackBarService.showSuccess(
                          // ignore: use_build_context_synchronously
                          context,
                          'Transação atualizada com sucesso!',
                        );
                        Modular.to.pop();
                      }
                    },
                    type: VRButtonType.primary,
                  ),
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }

  void _confirmSendTransaction(String id) async {
    await store.sendPendingTransaction(id);

    if (!mounted) return;

    if (store.errorMessage != null) {
      SnackBarService.showError(context, store.errorMessage!);
    } else {
      SnackBarService.showSuccess(context, 'Transação enviada com sucesso!');
    }
  }

  void _confirmDeleteTransaction(String id) async {
    await store.deletePendingTransaction(id);

    if (!mounted) return;

    if (store.errorMessage != null) {
      SnackBarService.showError(context, store.errorMessage!);
    } else {
      SnackBarService.showSuccess(context, 'Transação deletada com sucesso!');
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: const VRAppBar(),
      body: Padding(
        padding: const EdgeInsets.all(24),
        child: Center(
          child: ConstrainedBox(
            constraints: const BoxConstraints(maxWidth: 500),
            child: Observer(
              builder: (_) {
                if (store.isLoading) {
                  return const Center(child: CircularProgressIndicator());
                }

                if (store.pendingTransactions.isEmpty) {
                  return Column(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      const Text(
                        'Nenhuma transação pendente encontrada.',
                        style: TextStyle(fontSize: 18),
                        textAlign: TextAlign.center,
                      ),
                      const SizedBox(height: 260),
                      VRButton(
                        icon: Icons.arrow_back,
                        label: 'Voltar',
                        onTap: () => Modular.to.pop(),
                        type: VRButtonType.outlined,
                      ),
                    ],
                  );
                }

                return Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Expanded(
                      child: ListView.builder(
                        itemCount: store.pendingTransactions.length,
                        itemBuilder: (context, index) {
                          final transaction = store.pendingTransactions[index];
                          return TransactionCard(
                            description: transaction.description,
                            date: transaction.date,
                            originalValue: transaction.amountUsd,
                            actions: [
                              IconButton(
                                icon: const Icon(
                                  Icons.edit,
                                  color: Colors.blueAccent,
                                ),
                                onPressed: () =>
                                    _editTransaction(context, transaction.id),
                              ),
                              IconButton(
                                icon: const Icon(
                                  Icons.delete,
                                  color: Colors.redAccent,
                                ),
                                onPressed: () =>
                                    _confirmDeleteTransaction(transaction.id),
                              ),
                              IconButton(
                                icon: const Icon(
                                  Icons.cloud_upload,
                                  color: Colors.green,
                                ),
                                onPressed: () =>
                                    _confirmSendTransaction(transaction.id),
                              ),
                            ],
                          );
                        },
                      ),
                    ),
                    const SizedBox(height: 24),
                    VRButton(
                      icon: Icons.arrow_back,
                      label: 'Voltar',
                      onTap: () => Modular.to.pop(),
                      type: VRButtonType.outlined,
                    ),
                  ],
                );
              },
            ),
          ),
        ),
      ),
    );
  }
}
