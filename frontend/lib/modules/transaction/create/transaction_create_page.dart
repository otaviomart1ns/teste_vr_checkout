import 'package:flutter/material.dart';
import 'package:flutter_mobx/flutter_mobx.dart';
import 'package:flutter_modular/flutter_modular.dart';
import 'package:flutter_multi_formatter/formatters/formatter_utils.dart';
import 'package:frontend/core/services/snackbar_service.dart';
import 'package:frontend/core/widgets/app_bar.dart';
import 'package:frontend/core/widgets/button.dart';
import 'package:frontend/core/widgets/transaction_form_fields.dart';
import 'package:frontend/modules/transaction/create/store/transaction_create_store.dart';

class TransactionCreatePage extends StatefulWidget {
  const TransactionCreatePage({super.key});

  @override
  State<TransactionCreatePage> createState() => _TransactionCreatePageState();
}

class _TransactionCreatePageState extends State<TransactionCreatePage> {
  final _formKey = GlobalKey<FormState>();
  final _descriptionController = TextEditingController();
  final _valueController = TextEditingController();
  DateTime _selectedDate = DateTime.now();

  late final TransactionCreateStore store;

  @override
  void initState() {
    super.initState();
    store = Modular.get<TransactionCreateStore>();
  }

  @override
  void dispose() {
    _descriptionController.dispose();
    _valueController.dispose();
    super.dispose();
  }

  Future<void> _saveRemoteTransaction() async {
    if (!(_formKey.currentState?.validate() ?? false)) return;

    final cleanValue = toNumericString(
      _valueController.text,
      allowPeriod: true,
    );
    final double? amountUsd = double.tryParse(cleanValue);

    if (amountUsd == null) {
      if (!mounted) return;
      SnackBarService.showError(context, 'Valor inválido.');
      return;
    }

    await store.createTransactionRemote(
      description: _descriptionController.text.trim(),
      date: _selectedDate,
      amountUsd: amountUsd,
    );

    if (!mounted) return;

    if (store.errorMessage != null) {
      SnackBarService.showError(context, store.errorMessage!);
    } else {
      SnackBarService.showSuccess(context, 'Transação enviada com sucesso!');
      if (mounted) {
        _descriptionController.clear();
        _valueController.clear();
        setState(() {
          _selectedDate = DateTime.now();
        });
      }
    }
  }

  Future<void> _saveLocalTransaction() async {
    if (!(_formKey.currentState?.validate() ?? false)) return;

    final amount = double.tryParse(
      _valueController.text.replaceAll(',', '').replaceAll('\$', ''),
    );

    if (amount == null) {
      if (!mounted) return;
      SnackBarService.showError(context, 'Valor inválido.');
      return;
    }

    await store.createTransactionLocal(
      description: _descriptionController.text,
      date: _selectedDate,
      amountUsd: amount,
    );

    if (!mounted) return;
    SnackBarService.showSuccess(context, 'Transação salva localmente!');
    if (mounted) {
      _descriptionController.clear();
      _valueController.clear();
      setState(() {
        _selectedDate = DateTime.now();
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: const VRAppBar(),
      body: SingleChildScrollView(
        padding: const EdgeInsets.all(24),
        child: Center(
          child: ConstrainedBox(
            constraints: const BoxConstraints(maxWidth: 800),
            child: Observer(
              builder: (_) => Form(
                key: _formKey,
                child: Column(
                  children: [
                    const SizedBox(height: 160),
                    TransactionFormFields(
                      descriptionController: _descriptionController,
                      valueController: _valueController,
                      selectedDate: _selectedDate,
                      onDateChanged: (newDate) {
                        setState(() {
                          _selectedDate = newDate;
                        });
                      },
                    ),
                    const SizedBox(height: 32),
                    Row(
                      children: [
                        Expanded(
                          child: VRButton(
                            icon: Icons.arrow_back,
                            label: 'Voltar',
                            onTap: () => Modular.to.pop(),
                            type: VRButtonType.outlined,
                          ),
                        ),
                        const SizedBox(width: 16),
                        Expanded(
                          child: VRButton(
                            icon: Icons.save_alt,
                            label: 'Salvar Localmente',
                            onTap: () => _saveLocalTransaction(),
                            type: VRButtonType.outlined,
                          ),
                        ),
                        const SizedBox(width: 16),
                        Expanded(
                          child: VRButton(
                            icon: Icons.save,
                            label: 'Salvar',
                            onTap: store.isLoading
                                ? null
                                : () => _saveRemoteTransaction(),
                            type: VRButtonType.primary,
                          ),
                        ),
                      ],
                    ),
                    if (store.isLoading)
                      const Padding(
                        padding: EdgeInsets.only(top: 24),
                        child: CircularProgressIndicator(),
                      ),
                  ],
                ),
              ),
            ),
          ),
        ),
      ),
    );
  }
}
