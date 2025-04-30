import 'package:flutter/material.dart';
import 'package:flutter_mobx/flutter_mobx.dart';
import 'package:flutter_modular/flutter_modular.dart';
import 'package:frontend/core/utils/validators.dart';
import 'package:frontend/core/widgets/latest_transactions_card.dart';
import 'package:frontend/core/widgets/transaction_card.dart';
import 'package:frontend/core/widgets/app_bar.dart';
import 'package:frontend/core/widgets/button.dart';
import 'package:frontend/modules/transaction/view/store/transaction_view_store.dart';

class TransactionViewPage extends StatefulWidget {
  const TransactionViewPage({super.key});

  @override
  State<TransactionViewPage> createState() => _TransactionViewPageState();
}

class _TransactionViewPageState extends State<TransactionViewPage> {
  final _formKey = GlobalKey<FormState>();
  final _idController = TextEditingController();
  String? _selectedCurrency;
  late final TransactionViewStore store;

  bool hasConsulted = false;

  @override
  void initState() {
    super.initState();
    store = Modular.get<TransactionViewStore>();
    store.fetchCurrencies();
    store.fetchLatestTransactions();
  }

  @override
  void dispose() {
    _idController.dispose();
    super.dispose();
  }

  Future<void> _consultTransaction() async {
    if (!(_formKey.currentState?.validate() ?? false)) return;
    if (_selectedCurrency == null) {
      ScaffoldMessenger.of(
        context,
      ).showSnackBar(const SnackBar(content: Text('Selecione uma moeda.')));
      return;
    }
    await store.fetchTransaction(_idController.text.trim(), _selectedCurrency!);
    setState(() {
      hasConsulted = true;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: const VRAppBar(),
      body: SingleChildScrollView(
        padding: const EdgeInsets.all(24.0),
        child: Center(
          child: ConstrainedBox(
            constraints: const BoxConstraints(maxWidth: 800),
            child: Form(
              key: _formKey,
              child: Observer(
                builder: (_) => Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    const SizedBox(height: 120),
                    TextFormField(
                      controller: _idController,
                      decoration: const InputDecoration(
                        labelText: 'ID da Transação',
                        border: OutlineInputBorder(),
                      ),
                      validator: Validators.validateTransactionId,
                    ),
                    const SizedBox(height: 16),
                    DropdownButtonFormField<String>(
                      value: _selectedCurrency,
                      decoration: const InputDecoration(
                        labelText: 'Moeda de Destino',
                        border: OutlineInputBorder(),
                      ),
                      items: store.currencies.map((currency) {
                        return DropdownMenuItem(
                          value: currency,
                          child: Text(currency),
                        );
                      }).toList(),
                      onChanged: (value) {
                        setState(() {
                          _selectedCurrency = value;
                        });
                      },
                      validator: (value) =>
                          value == null ? 'Escolha uma moeda' : null,
                    ),
                    const SizedBox(height: 34),
                    if (store.isLoading)
                      const Center(child: CircularProgressIndicator())
                    else if (hasConsulted && store.transaction != null)
                      TransactionCard(
                        description: store.transaction!['description'],
                        date: DateTime.parse(store.transaction!['date']),
                        originalValue: (store.transaction!['amount_usd'] as num)
                            .toDouble(),
                        convertedValue:
                            (store.transaction!['amount_converted'] as num)
                                .toDouble(),
                        exchangeRate:
                            (store.transaction!['exchange_rate'] as num)
                                .toDouble(),
                      )
                    else if (!hasConsulted)
                      LatestTransactionsWidget(
                        transactions: store.latestTransactions,
                      ),
                    const SizedBox(height: 34),
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
                            icon: Icons.search,
                            label: 'Consultar',
                            onTap: _consultTransaction,
                            type: VRButtonType.primary,
                          ),
                        ),
                      ],
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
